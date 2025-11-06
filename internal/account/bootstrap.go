package account

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/account/adapter"
	"shikposh-backend/internal/account/entryporint/handler"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"shikposh-backend/internal/account/service_layer/event_handler"
	"shikposh-backend/pkg/framework/infrastructure/logging"
	commandeventhandler "shikposh-backend/pkg/framework/service_layer/command_event_handler"
	"shikposh-backend/pkg/framework/service_layer/messagebus"
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config, logger logging.Logger, bus messagebus.MessageBus) error {
	uow := unit_of_work.New(db)

	ag, err := adapter.NewAvatarGenerator(AssetsFS)
	if err != nil {
		logging.Error("Failed to initialize avatar generator").WithError(err).Log()
		return err
	}

	userHandler := command_handler.NewUserHandler(uow, cfg)
	userEventHandler := event_handler.NewUserEventHandler(uow)
	userController := handler.NewUserController(bus, ag)
	userController.RegisterRoutes(router)

	bus.AddHandler(
		commandeventhandler.NewCommandHandlerWithResult(userHandler.RegisterHandler),
		commandeventhandler.NewCommandHandlerWithResult(userHandler.LoginHandler),
		commandeventhandler.NewCommandHandler(userHandler.LogoutHandler),
	)
	bus.AddHandlerEvent(
		commandeventhandler.NewEventHandler(userEventHandler.RegisterEvent),
	)

	return nil
}
