package query

func (u *update) WithNot(k string, v interface{}) Update {
	if u.nots == nil {
		u.nots = make([]KeyValue, 0)
	}
	u.nots = append(u.nots, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetNots() []KeyValue {
	return u.nots
}
