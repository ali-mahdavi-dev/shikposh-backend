package command_handler

import (
	unitofwork "shikposh-backend/internal/unit_of_work"

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

type WishlistCommandHandler struct {
	uow unitofwork.PGUnitOfWork
}

func NewWishlistCommandHandler(uow unitofwork.PGUnitOfWork) *WishlistCommandHandler {
	return &WishlistCommandHandler{uow: uow}
}
