package query

type Operator byte

const (
	None Operator = iota
	Equal
	NotEqual
	Lower
	LowerEqual
	Greater
	GreaterEqual
	BitwiseIs
	BitwiseIsNot
	BitwiseAndEqual
	BitwiseAndNotEqual
	In
	Is
	IsNot
	NotIn
)
