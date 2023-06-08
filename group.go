package query

type Group interface {
	Query
	GroupBy(fields ...string) Group
	GetGroupBy() []string
}

type group struct {
	Query
	fields []string
}

func NewGroup(base Query) Group {
	return &group{
		Query: base,
	}
}

func (s *group) GetSub() Query {
	return s.Query
}

func (s *group) GetGroup() Group {
	return s
}

func (s *group) GroupBy(fields ...string) Group {
	s.fields = fields
	return s
}

func (s *group) GetGroupBy() []string {
	return s.fields
}
