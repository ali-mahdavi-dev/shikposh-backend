package command_handler

import (
	"shikposh-backend/config"
	"shikposh-backend/internal/account/adapter"
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type OtpHandler struct {
	uow        unitofwork.PGUnitOfWork
	cfg        *config.Config
	otpService *adapter.OtpService
}

func NewOtpHandler(uow unitofwork.PGUnitOfWork, cfg *config.Config, otpService *adapter.OtpService) *OtpHandler {
	return &OtpHandler{
		uow:        uow,
		cfg:        cfg,
		otpService: otpService,
	}
}

type UserHandler struct {
	uow unitofwork.PGUnitOfWork
	cfg *config.Config
}

func NewUserHandler(uow unitofwork.PGUnitOfWork, cfg *config.Config) *UserHandler {
	return &UserHandler{uow: uow, cfg: cfg}
}
