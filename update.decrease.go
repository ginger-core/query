package query

func (u *update) WithDecrease(k string, v interface{}) Update {
	if u.decreases == nil {
		u.decreases = make([]KeyValue, 0)
	}
	u.decreases = append(u.decreases, KeyValue{Key: k, Value: v})
	return u
}

func (u *update) GetDecreases() []KeyValue {
	return u.decreases
}
