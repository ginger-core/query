package query

type ModelHandlerFunc func() any

type ModelQuery interface {
	Query
	WithModelHandlerFunc(ModelHandlerFunc) ModelQuery
	GetModelHandlerFunc() ModelHandlerFunc
	GetModel() any
}

type modelQuery struct {
	Query
	modelHandlerFunc ModelHandlerFunc
}

func NewModelQuery(base Query) ModelQuery {
	return &modelQuery{
		Query: base,
	}
}

func (q *modelQuery) GetSub() Query {
	return q.Query
}

func (q *modelQuery) WithModelHandlerFunc(handlerFunc ModelHandlerFunc) ModelQuery {
	q.modelHandlerFunc = handlerFunc
	return q
}

func (q *modelQuery) GetModelHandlerFunc() ModelHandlerFunc {
	return q.modelHandlerFunc
}

func (q *modelQuery) GetModel() any {
	if q.modelHandlerFunc == nil {
		return nil
	}
	return q.modelHandlerFunc()
}
