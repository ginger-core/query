package query

func (u *update) WithIncrease(k string, v interface{}) Update {
	if u.increases == nil {
		u.increases = make([]KeyValue, 0)
	}
	u.increases = append(u.increases, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetIncreases() []KeyValue {
	return u.increases
}
