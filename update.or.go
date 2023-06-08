package query

func (u *update) WithOr(k string, v interface{}) Update {
	if u.ors == nil {
		u.ors = make([]KeyValue, 0)
	}
	u.ors = append(u.ors, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetOrs() []KeyValue {
	return u.ors
}
