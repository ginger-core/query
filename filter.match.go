package query

type Match struct {
	Key      string
	Value    any
	Operator Operator
	// Extra is being used for specific filters like bitwise filters
	// which needs to operate on a value and compare result with
	Extra []any

	customHandle CustomHandlerFunc
}

func (f *filter) WithMatch(m *Match) Filter {
	f.matchs = append(f.matchs, m)
	return f
}

func (f *filter) GetMatchs() []*Match {
	return f.matchs
}
