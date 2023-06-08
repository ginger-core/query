package query

type KeyValue struct {
	Key   string
	Value interface{}
}

type Update interface {
	Query
	GetUpdate() Update

	WithSet(k string, v interface{}) Update
	GetSets() []KeyValue

	WithIncrease(k string, v interface{}) Update
	GetIncreases() []KeyValue

	WithDecrease(k string, v interface{}) Update
	GetDecreases() []KeyValue

	WithAnd(k string, v interface{}) Update
	GetAnds() []KeyValue

	WithOr(k string, v interface{}) Update
	GetOrs() []KeyValue

	WithNot(k string, v interface{}) Update
	GetNots() []KeyValue
}

type update struct {
	Query
	sets      []KeyValue
	increases []KeyValue
	decreases []KeyValue
	ands      []KeyValue
	ors       []KeyValue
	nots      []KeyValue
}

func NewUpdate(base Query) Update {
	return &update{
		Query: base,
	}
}

func (u *update) GetSub() Query {
	return u.Query
}

func (u *update) GetUpdate() Update {
	return u
}
