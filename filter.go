package query

type Filter interface {
	Query

	WithBase(q Query) Filter

	WithId(id any) Filter
	GetId() any

	WithMatch(m *Match) Filter
	GetMatchs() []*Match

	WithAnd(filters ...Filter) Filter
	WithOr(filters ...Filter) Filter

	GetOperations() []FilterOperation
}

type filter struct {
	Query
	id         any
	matchs     []*Match
	operations []FilterOperation
}

func NewFilter(base Query) Filter {
	return &filter{
		Query:  base,
		matchs: make([]*Match, 0),
	}
}

func (f *filter) WithBase(q Query) Filter {
	f.Query = q
	return f
}

func (f *filter) GetSub() Query {
	return f.Query
}

func (f *filter) GetFilter() Filter {
	return f
}
