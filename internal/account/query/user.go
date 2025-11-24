package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type UserQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewUserQueryHandler(uow unitofwork.PGUnitOfWork) *UserQueryHandler {
	return &UserQueryHandler{uow: uow}
}
