package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type CategoryQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewCategoryQueryHandler(uow unitofwork.PGUnitOfWork) *CategoryQueryHandler {
	return &CategoryQueryHandler{uow: uow}
}
