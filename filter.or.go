package query

func (f *filter) WithOr(filters ...Filter) Filter {
	for _, filter := range filters {
		f.operations = append(f.operations,
			&filterOperation{
				filter: filter,
				op:     OperationOr,
			})
	}
	return f
}
