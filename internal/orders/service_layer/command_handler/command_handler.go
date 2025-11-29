package command_handler

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type OrderCommandHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewOrderCommandHandler(uow unitofwork.PGUnitOfWork) *OrderCommandHandler {
	return &OrderCommandHandler{
		uow: uow,
	}
}

func (h *OrderCommandHandler) GetUOW() unitofwork.PGUnitOfWork {
	return h.uow
}
