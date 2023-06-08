package query

type ModelsHandlerFunc func() any

type ModelsQuery interface {
	Query
	WithModelsHandlerFunc(ModelsHandlerFunc) ModelsQuery
	GetModelsHandlerFunc() ModelsHandlerFunc
	GetModels() any
}

type modelsQuery struct {
	Query
	modelsHandlerFunc ModelsHandlerFunc
}

func NewModelsQuery(base Query) ModelsQuery {
	return &modelsQuery{
		Query: base,
	}
}

func (q *modelsQuery) GetSub() Query {
	return q.Query
}

func (q *modelsQuery) WithModelsHandlerFunc(handlerFunc ModelsHandlerFunc) ModelsQuery {
	q.modelsHandlerFunc = handlerFunc
	return q
}

func (q *modelsQuery) GetModelsHandlerFunc() ModelsHandlerFunc {
	return q.modelsHandlerFunc
}

func (q *modelsQuery) GetModels() any {
	if q.modelsHandlerFunc == nil {
		return nil
	}
	return q.modelsHandlerFunc()
}
