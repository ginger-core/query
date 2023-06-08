package query

func (f *filter) WithId(id any) Filter {
	f.id = id
	return f.WithMatch(&Match{
		Key:      "id",
		Operator: Equal,
		Value:    id,
	})
}

func (f *filter) GetId() any {
	return f.id
}
