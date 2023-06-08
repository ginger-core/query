package query

type Operation int

const (
	OperationAnd Operation = iota
	OperationOr
)

type FilterOperation interface {
	GetFilter() Filter
	GetOperation() Operation
}

type filterOperation struct {
	filter Filter
	op     Operation
}

func (fo *filterOperation) GetFilter() Filter {
	return fo.filter
}

func (fo *filterOperation) GetOperation() Operation {
	return fo.op
}

func (f *filter) GetOperations() []FilterOperation {
	return f.operations
}
