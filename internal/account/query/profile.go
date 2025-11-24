package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type ProfileQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewProfileQueryHandler(uow unitofwork.PGUnitOfWork) *ProfileQueryHandler {
	return &ProfileQueryHandler{uow: uow}
}
