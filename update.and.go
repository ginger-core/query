package query

func (u *update) WithAnd(k string, v interface{}) Update {
	if u.ands == nil {
		u.ands = make([]KeyValue, 0)
	}
	u.ands = append(u.ands, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetAnds() []KeyValue {
	return u.ands
}
