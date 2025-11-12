package account

import (
	"shikposh-backend/config"
	accountadapter "shikposh-backend/internal/account/adapter"
	"shikposh-backend/internal/account/entrypoint"
	"shikposh-backend/internal/account/entrypoint/handler"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"shikposh-backend/internal/account/service_layer/event_handler"
	"shikposh-backend/pkg/framework/adapter"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	commandeventhandler "shikposh-backend/pkg/framework/service_layer/command_event_handler"
	commandmiddleware "shikposh-backend/pkg/framework/service_layer/command_event_handler/command_middleware"
	"shikposh-backend/pkg/framework/service_layer/messagebus"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config) error {
	// Create event channel and unit of work for this module
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unit_of_work.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	ag, err := accountadapter.NewAvatarGenerator(AssetsFS)
	if err != nil {
		logging.Error("Failed to initialize avatar generator").WithError(err).Log()
		return err
	}

	userHandler := command_handler.NewUserHandler(uow, cfg)
	userEventHandler := event_handler.NewUserEventHandler(uow)
	userController := handler.NewUserController(bus, ag, userHandler)

	entrypoint.NewAccountRouter(router, entrypoint.UserManagementRouter{
		User: userController,
	})

	// register command middlewares
	bus.AddCommandMiddleware(
		commandmiddleware.Logging(),
	)

	// register command handlers
	bus.AddCommandHandler(
		commandeventhandler.NewCommandHandler(userHandler.RegisterHandler),
		commandeventhandler.NewCommandHandler(userHandler.LogoutHandler),
	)

	// register event handlers
	bus.AddEventHandler(
		commandeventhandler.NewEventHandler(userEventHandler.RegisterEvent),
	)

	return nil
}
