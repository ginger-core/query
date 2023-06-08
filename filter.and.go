package query

func (f *filter) WithAnd(filters ...Filter) Filter {
	for _, filter := range filters {
		f.operations = append(f.operations,
			&filterOperation{
				filter: filter,
				op:     OperationAnd,
			})
	}
	return f
}
