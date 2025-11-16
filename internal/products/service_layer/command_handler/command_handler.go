package command_handler

import (
	"shikposh-backend/internal/unit_of_work"

	"github.com/gosimple/slug"
)

func GenerateSlug(name string) string {
	generatedSlug := slug.Make(name)

	return generatedSlug
}

type ProductCommandHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewProductCommandHandler(uow unitofwork.PGUnitOfWork) *ProductCommandHandler {
	return &ProductCommandHandler{uow: uow}
}

type ReviewCommandHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewReviewCommandHandler(uow unitofwork.PGUnitOfWork) *ReviewCommandHandler {
	return &ReviewCommandHandler{uow: uow}
}
