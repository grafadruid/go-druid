package aggregation

// First aggregations using generics for numeric types

type NumericFirst[T any] struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func (n *NumericFirst[T]) SetName(name string) *NumericFirst[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericFirst[T]) SetFieldName(fieldName string) *NumericFirst[T] {
	n.FieldName = fieldName
	return n
}

func (n *NumericFirst[T]) SetTimeColumn(timeColumn string) *NumericFirst[T] {
	n.TimeColumn = timeColumn
	return n
}

// Type aliases for backward compatibility
type (
	DoubleFirst = NumericFirst[float64]
	FloatFirst  = NumericFirst[float32]
	LongFirst   = NumericFirst[int64]
)

// Constructor functions
func NewDoubleFirst() *DoubleFirst {
	d := &DoubleFirst{}
	d.SetType("doubleFirst")
	return d
}

func NewFloatFirst() *FloatFirst {
	f := &FloatFirst{}
	f.SetType("floatFirst")
	return f
}

func NewLongFirst() *LongFirst {
	l := &LongFirst{}
	l.SetType("longFirst")
	return l
}
