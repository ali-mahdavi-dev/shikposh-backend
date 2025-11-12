package messagebus

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"shikposh-backend/pkg/framework/adapter"
	apperrors "shikposh-backend/pkg/framework/errors"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	commandeventhandler "shikposh-backend/pkg/framework/service_layer/command_event_handler"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"
)

type MessageBus interface {
	AddHandler(handlers ...commandeventhandler.CommandHandler) error
	AddHandlerEvent(handlers ...commandeventhandler.EventHandler) error
	Handle(ctx context.Context, cmd any) error
	Uow() unit_of_work.PGUnitOfWork
	Shutdown(ctx context.Context) error
	EventChannel() chan<- adapter.EventWithWaitGroup
}

type messageBus struct {
	handledCommands map[any]commandeventhandler.CommandHandler
	handledEvent    map[any]commandeventhandler.EventHandler
	uow             unit_of_work.PGUnitOfWork
	eventCh         chan adapter.EventWithWaitGroup
	shutdownCh      chan struct{}
	wg              sync.WaitGroup
	mu              sync.RWMutex
	closed          bool
}

func NewMessageBus(uow unit_of_work.PGUnitOfWork, eventCh chan adapter.EventWithWaitGroup) MessageBus {
	bus := &messageBus{
		handledCommands: make(map[any]commandeventhandler.CommandHandler),
		handledEvent:    make(map[any]commandeventhandler.EventHandler),
		uow:             uow,
		eventCh:         eventCh,
		shutdownCh:      make(chan struct{}),
		closed:          false,
	}

	logging.Info("Message bus initialized").
		WithInt("event_channel_capacity", 100).
		Log()

	// start event handler worker pool
	// Using a single goroutine to process events sequentially to avoid race conditions
	// Events are processed one at a time to ensure proper ordering and prevent concurrency issues
	bus.wg.Add(1)
	go func(mb *messageBus, evCh chan adapter.EventWithWaitGroup) {
		defer mb.wg.Done()
		defer close(evCh)
		for eventWrapper := range evCh {
			eventCtx := eventWrapper.Ctx
			if eventCtx == nil {
				eventCtx = context.Background()
			}
			if err := mb.HandleEvent(eventCtx, eventWrapper.Event); err != nil {
				logging.Error("Failed to handle event").WithError(err).Log()
			}
			// Signal that event is done being handled
			if eventWrapper.Wg != nil {
				eventWrapper.Wg.Done()
			}
		}
	}(bus, bus.eventCh)

	return bus
}

func (m *messageBus) Uow() unit_of_work.PGUnitOfWork {
	return m.uow
}

func (m *messageBus) AddHandler(handlers ...commandeventhandler.CommandHandler) error {
	for _, handler := range handlers {
		cmdName := reflect.TypeOf(handler.NewCommand()).String()
		if _, ok := m.handledCommands[cmdName]; ok {
			return apperrors.Conflict("", fmt.Sprintf("command handler for %s already exists", cmdName))
		}
		m.handledCommands[cmdName] = handler
		logging.Info("Command handler registered").
			WithAny("command_name", cmdName).
			Log()
	}

	return nil
}

func (m *messageBus) AddHandlerEvent(handlers ...commandeventhandler.EventHandler) error {
	for _, handler := range handlers {
		eventName := reflect.TypeOf(handler.NewEvent()).String()
		if _, ok := m.handledEvent[eventName]; ok {
			return apperrors.Conflict("", fmt.Sprintf("event handler for %s already exists", eventName))
		}
		m.handledEvent[eventName] = handler
		logging.Info("Event handler registered").
			WithAny("event_name", eventName).
			Log()
	}

	return nil
}

func (m *messageBus) AddEvent(handlers ...commandeventhandler.EventHandler) error {
	for _, handler := range handlers {
		eventName := reflect.TypeOf(handler.NewEvent()).String()
		if _, ok := m.handledEvent[eventName]; ok {
			return apperrors.Conflict("", fmt.Sprintf("event handler for %s already exists", eventName))
		}
		m.handledEvent[eventName] = handler
	}

	return nil
}

func (m *messageBus) Handle(ctx context.Context, cmd any) error {
	cmdName := reflect.TypeOf(cmd).String()

	logging.Info("Handling command").
		WithAny("command_name", cmdName).
		Log()

	handler, ok := m.handledCommands[cmdName]
	if !ok {
		err := apperrors.NotFound("", fmt.Sprintf("command handler for %s not found", cmdName))
		logging.Error("Command handler not found").
			WithAny("command_name", cmdName).
			WithError(err).
			Log()
		return err
	}

	err := handler.Handle(ctx, cmd)
	if err != nil {
		logging.Error("Command handler failed").
			WithAny("command_name", cmdName).
			WithError(err).
			Log()
		return err
	}

	logging.Info("Command handled successfully").
		WithAny("command_name", cmdName).
		Log()

	return nil
}

func (m *messageBus) HandleEvent(ctx context.Context, event any) error {
	eventName := reflect.TypeOf(event).String()

	logging.Info("Handling event").
		WithAny("event_name", eventName).
		Log()

	if _, ok := m.handledEvent[eventName]; !ok {
		err := apperrors.NotFound("", fmt.Sprintf("event handler for %s not found", eventName))
		logging.Error("Event handler not found").
			WithAny("event_name", eventName).
			WithError(err).
			Log()
		return err
	}

	err := m.handledEvent[eventName].Handle(ctx, event)
	if err != nil {
		logging.Error("Event handler failed").
			WithAny("event_name", eventName).
			WithError(err).
			Log()
		return err
	}

	logging.Info("Event handled successfully").
		WithAny("event_name", eventName).
		Log()

	return nil
}

// Shutdown gracefully shuts down the message bus.
// It stops accepting new events, processes remaining events, and closes the event channel.
func (m *messageBus) Shutdown(ctx context.Context) error {
	m.mu.Lock()
	if m.closed {
		m.mu.Unlock()
		return nil
	}
	m.closed = true
	m.mu.Unlock()

	// Signal shutdown to the event handler goroutine
	close(m.shutdownCh)

	// Wait for all events to be processed or context timeout
	done := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logging.Info("Message bus shutdown completed successfully").Log()
		return nil
	case <-ctx.Done():
		logging.Warn("Message bus shutdown timed out").WithError(ctx.Err()).Log()
		return ctx.Err()
	}
}

// EventChannel returns the event channel for use by unit of work
func (m *messageBus) EventChannel() chan<- adapter.EventWithWaitGroup {
	return m.eventCh
}
