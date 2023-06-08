package query

func (u *update) WithSet(k string, v interface{}) Update {
	if u.sets == nil {
		u.sets = make([]KeyValue, 0)
	}
	u.sets = append(u.sets, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetSets() []KeyValue {
	return u.sets
}
