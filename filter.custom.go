package query

import "github.com/ginger-core/errors"

type CustomHandlerFunc func(q Query, match *Match) (FilterResult, errors.Error)

func (m *Match) WithCustomHandle(_f CustomHandlerFunc) *Match {
	m.customHandle = _f
	return m
}

func (m *Match) HasCustomHandle() bool {
	return m.customHandle != nil
}

func (m *Match) HandleCustom(q Query) (FilterResult, errors.Error) {
	return m.customHandle(q, m)
}

type FilterResult interface {
	WithData(d any) FilterResult
	GetData() any
}

type filterResult struct {
	data any
}

func NewResult() FilterResult {
	return new(filterResult)
}

func (r *filterResult) WithData(d any) FilterResult {
	r.data = d
	return r
}

func (r *filterResult) GetData() any {
	return r.data
}
