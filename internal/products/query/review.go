package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"
)

type ReviewQueryHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewReviewQueryHandler(uow unitofwork.PGUnitOfWork) *ReviewQueryHandler {
	return &ReviewQueryHandler{uow: uow}
}
