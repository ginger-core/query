package query

type Pagination interface {
	Query
	WithSize(size int) Pagination
	GetSize() int

	WithPage(page int) Pagination
	GetPage() int
}

type pagination struct {
	Query
	size int
	page int
}

func NewPagination(base Query) Pagination {
	return &pagination{
		Query: base,
	}
}

func (p *pagination) GetSub() Query {
	return p.Query
}

func (p *pagination) GetPagination() Pagination {
	return p
}

func (p *pagination) WithSize(size int) Pagination {
	p.size = size
	return p
}

func (p *pagination) GetSize() int {
	return p.size
}

func (p *pagination) WithPage(page int) Pagination {
	p.page = page
	return p
}

func (p *pagination) GetPage() int {
	return p.page
}
