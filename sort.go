package query

type Sort interface {
	Query
	GetSortBy() string
	SortBy(field string) Sort
	Asc() Sort
	Desc() Sort
	IsAsc() bool
}

type sort struct {
	Query
	field string
	desc  bool
}

func NewSort(base Query) Sort {
	return &sort{
		Query: base,
	}
}

func (s *sort) GetSub() Query {
	return s.Query
}

func (s *sort) GetSort() Sort {
	return s
}

func (s *sort) SortBy(field string) Sort {
	s.field = field
	return s
}

func (s *sort) GetSortBy() string {
	return s.field
}

func (s *sort) Asc() Sort {
	s.desc = false
	return s
}

func (s *sort) Desc() Sort {
	s.desc = true
	return s
}

func (s *sort) IsAsc() bool {
	return !s.desc
}

type Sorts interface {
	Query
	Clone() Sorts
	WithSort(s Sort) Sorts
	NextSort() Sort
}

type sorts struct {
	Query
	items []Sort
}

func NewSorts(base Query) Sorts {
	return &sorts{
		Query: base,
	}
}

func (s *sorts) Clone() Sorts {
	s2 := &sorts{
		Query: s.Query,
		items: s.items,
	}
	return s2
}

func (s *sorts) GetSorts() Sorts {
	return s
}

func (s *sorts) WithSort(sort Sort) Sorts {
	s.items = append(s.items, sort)
	return s
}

func (s *sorts) NextSort() Sort {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}
