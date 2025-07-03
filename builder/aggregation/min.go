package aggregation

// Min aggregations using generics for numeric types

type NumericMin[T any] struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func (n *NumericMin[T]) SetName(name string) *NumericMin[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericMin[T]) SetFieldName(fieldName string) *NumericMin[T] {
	n.FieldName = fieldName
	return n
}

func (n *NumericMin[T]) SetExpression(expression string) *NumericMin[T] {
	n.Expression = expression
	return n
}

// Type aliases for backward compatibility
type (
	DoubleMin = NumericMin[float64]
	FloatMin  = NumericMin[float32]
	LongMin   = NumericMin[int64]
)

// Constructor functions
func NewDoubleMin() *DoubleMin {
	d := &DoubleMin{}
	d.SetType("doubleMin")
	return d
}

func NewFloatMin() *FloatMin {
	f := &FloatMin{}
	f.SetType("floatMin")
	return f
}

func NewLongMin() *LongMin {
	l := &LongMin{}
	l.SetType("longMin")
	return l
}
