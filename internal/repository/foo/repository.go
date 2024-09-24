package foo

import (
	"github.com/jmoiron/sqlx"

	"github.com/north70/go-template/internal/repository"
)

const (
	table      = "foos"
	columnID   = "id"
	columnName = "name"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository.FooRepository {
	return &Repository{db: db}
}
