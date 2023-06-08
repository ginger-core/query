package query

import "context"

func (q *query) WithContext(ctx context.Context) Query {
	q.ctx = ctx
	return q
}

func (q *query) GetContext() context.Context {
	return q.ctx
}
