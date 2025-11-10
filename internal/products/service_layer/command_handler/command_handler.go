package command_handler

import (
	"shikposh-backend/pkg/framework/service_layer/unit_of_work"

	"github.com/gosimple/slug"
)

func GenerateSlug(name string) string {
	generatedSlug := slug.Make(name)

	return generatedSlug
}

type ProductCommandHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewProductCommandHandler(uow unit_of_work.PGUnitOfWork) *ProductCommandHandler {
	return &ProductCommandHandler{uow: uow}
}

type ReviewCommandHandler struct {
	uow unit_of_work.PGUnitOfWork
}

func NewReviewCommandHandler(uow unit_of_work.PGUnitOfWork) *ReviewCommandHandler {
	return &ReviewCommandHandler{uow: uow}
}
