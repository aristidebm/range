package utils

type BoundType string

const (
	LPAREN   BoundType = "("
	RPAREN             = ")"
	LBRACKET           = "["
	RBRACKET           = "]"
)

type Bound struct {
	Value float64
	Type  BoundType
}

func (b Bound) Equal(other Bound) bool {
	return b.Value == other.Value && b.Type == other.Type
}
