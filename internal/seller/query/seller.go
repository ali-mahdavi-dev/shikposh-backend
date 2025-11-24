package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type SellerQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewSellerQueryHandler(uow unitofwork.PGUnitOfWork) *SellerQueryHandler {
	return &SellerQueryHandler{
		uow: uow,
	}
}
