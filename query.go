package query

import (
	"context"

	"github.com/ginger-core/errors"
)

type Query interface {
	WithSub(q Query) Query
	GetSub() Query

	SetDB(db any)
	GetDB() any

	WithContext(ctx context.Context) Query
	GetContext() context.Context

	GetFilter() Filter
	GetSort() Sort
	GetSorts() Sorts
	GetPagination() Pagination
	GetGroup() Group

	GetUpdate() Update

	SetError(err errors.Error)
	GetError() errors.Error

	SetTag(key string, value any)
	GetTag(key string) any
	DeleteTag(key string)
}

type query struct {
	Query
	ctx context.Context
	err errors.Error

	tags map[string]any
}

func New(ctx context.Context) Query {
	return &query{
		ctx: ctx,
	}
}

func (q *query) WithSub(sub Query) Query {
	if q.Query != nil {
		return q.Query.WithSub(sub)
	}
	q.Query = sub
	return q
}

func (q *query) GetSub() Query {
	return q.Query
}

func (q *query) SetDB(db any) {
}

func (q *query) GetDB() any {
	return nil
}

func (q *query) GetFilter() Filter {
	return nil
}

func (q *query) GetSort() Sort {
	return nil
}

func (q *query) GetSorts() Sorts {
	return nil
}

func (q *query) GetPagination() Pagination {
	return nil
}

func (q *query) GetUpdate() Update {
	return nil
}

func (q *query) GetGroup() Group {
	return nil
}

func (q *query) SetError(err errors.Error) {
	q.err = err
}

func (q *query) GetError() errors.Error {
	return q.err
}

func (q *query) SetTag(key string, value any) {
	if q.tags == nil {
		q.tags = make(map[string]any)
	}
	q.tags[key] = value
}

func (q *query) GetTag(key string) any {
	return q.tags[key]
}

func (q *query) DeleteTag(key string) {
	delete(q.tags, key)
}
