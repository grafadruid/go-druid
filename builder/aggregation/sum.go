package aggregation

// Sum aggregations using generics for numeric types

type NumericSum[T any] struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	Expression string `json:"expression,omitempty"`
}

func (n *NumericSum[T]) SetName(name string) *NumericSum[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericSum[T]) SetFieldName(fieldName string) *NumericSum[T] {
	n.FieldName = fieldName
	return n
}

func (n *NumericSum[T]) SetExpression(expression string) *NumericSum[T] {
	n.Expression = expression
	return n
}

// Type aliases for backward compatibility
type (
	DoubleSum = NumericSum[float64]
	FloatSum  = NumericSum[float32]
	LongSum   = NumericSum[int64]
)

// Constructor functions
func NewDoubleSum() *DoubleSum {
	d := &DoubleSum{}
	d.SetType("doubleSum")
	return d
}

func NewFloatSum() *FloatSum {
	f := &FloatSum{}
	f.SetType("floatSum")
	return f
}

func NewLongSum() *LongSum {
	l := &LongSum{}
	l.SetType("longSum")
	return l
}
