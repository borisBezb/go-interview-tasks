package repo

import (
	"context"

	"github.com/doug-martin/goqu/v9"
)

type Repo[T any] interface {
	GetList(context.Context) ([]T, error)
}

func NewRepo[T any](db *goqu.Database, tableName string) Repo[T] {
	return &repo[T]{
		db:        db,
		tableName: tableName,
	}
}

type repo[T any] struct {
	db        *goqu.Database
	tableName string
}

func (r repo[T]) GetList(ctx context.Context) ([]T, error) {
	var res []T
	err := r.db.From(r.tableName).
		ScanStructs(&res)

	return res, err
}
