package account

import (
	"shikposh-backend/config"
	accountadapter "shikposh-backend/internal/account/adapter"
	accountphrases "shikposh-backend/internal/account/adapter/phrases"
	"shikposh-backend/internal/account/entrypoint"
	"shikposh-backend/internal/account/entrypoint/handler"
	"shikposh-backend/internal/account/service_layer/command_handler"
	"shikposh-backend/internal/account/service_layer/event_handler"
	unitofwork "shikposh-backend/internal/unit_of_work"

	"github.com/ali-mahdavi-dev/shikposh-framework/adapter"
	"github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/logging"
	redisx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/redisx"
	commandeventhandler "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler"
	commandmiddleware "github.com/ali-mahdavi-dev/shikposh-framework/service_layer/command_event_handler/command_middleware"
	"github.com/ali-mahdavi-dev/shikposh-framework/service_layer/messagebus"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(router fiber.Router, db *gorm.DB, cfg *config.Config, redis redisx.Connection) error {
	// Register account module error phrases
	accountphrases.RegisterAccountPhrases()
	// Create event channel and unit of work for this module
	eventCh := make(chan adapter.EventWithWaitGroup, 100)
	uow := unitofwork.New(db, eventCh)
	bus := messagebus.NewMessageBus(uow, eventCh)

	ag, err := accountadapter.NewAvatarGenerator(AssetsFS)
	if err != nil {
		logging.Error("Failed to initialize avatar generator").WithError(err).Log()
		return err
	}

	// Initialize OTP service
	otpService := accountadapter.NewOtpService(redis, cfg)

	userHandler := command_handler.NewUserHandler(uow, cfg)
	otpHandler := command_handler.NewOtpHandler(uow, cfg, otpService)
	userEventHandler := event_handler.NewUserEventHandler(uow)
	userController := handler.NewUserController(bus, ag, userHandler, otpHandler)

	entrypoint.NewAccountRouter(router, entrypoint.UserManagementRouter{
		User: userController,
	})

	// register command middlewares
	bus.AddCommandMiddleware(
		commandmiddleware.Logging(),
	)

	// register command handlers
	bus.AddCommandHandler(
		commandeventhandler.NewCommandHandler(userHandler.LogoutHandler),
		commandeventhandler.NewCommandHandler(otpHandler.SendOtpHandler),
	)

	// register event handlers
	bus.AddEventHandler(
		commandeventhandler.NewEventHandler(userEventHandler.RegisterEvent),
	)

	return nil
}
