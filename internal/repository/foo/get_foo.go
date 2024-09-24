package foo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/north70/go-template/internal/domain"
	"github.com/north70/go-template/internal/repository"
)

func (r *Repository) GetFoo(ctx context.Context, id string) (*domain.Foo, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(columnID, columnName).
		From(table).
		Where(sq.Eq{columnID: id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	row := r.db.QueryRowxContext(ctx, query, args...)
	if err = row.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var foo domain.Foo
	if err = row.StructScan(&foo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &foo, nil
}
