package query

import (
	unitofwork "shikposh-backend/internal/unit_of_work"

	elasticsearchx "github.com/ali-mahdavi-dev/shikposh-framework/infrastructure/elasticsearch"
)

type ProductQueryHandler struct {
	uow           unitofwork.PGUnitOfWork
	elasticsearch elasticsearchx.Connection
	indexName     string
}

func NewProductQueryHandler(uow unitofwork.PGUnitOfWork, elasticsearch elasticsearchx.Connection) *ProductQueryHandler {
	return &ProductQueryHandler{
		uow:           uow,
		elasticsearch: elasticsearch,
		indexName:     "products",
	}
}
