package aggregation

// Max aggregations using generics for numeric types

type NumericMax[T any] struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func (n *NumericMax[T]) SetName(name string) *NumericMax[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericMax[T]) SetFieldName(fieldName string) *NumericMax[T] {
	n.FieldName = fieldName
	return n
}

func (n *NumericMax[T]) SetExpression(expression string) *NumericMax[T] {
	n.Expression = expression
	return n
}

// Type aliases for backward compatibility
type (
	DoubleMax = NumericMax[float64]
	FloatMax  = NumericMax[float32]
	LongMax   = NumericMax[int64]
)

// Constructor functions
func NewDoubleMax() *DoubleMax {
	d := &DoubleMax{}
	d.SetType("doubleMax")
	return d
}

func NewFloatMax() *FloatMax {
	f := &FloatMax{}
	f.SetType("floatMax")
	return f
}

func NewLongMax() *LongMax {
	l := &LongMax{}
	l.SetType("longMax")
	return l
}
