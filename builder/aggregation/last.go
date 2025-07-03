package aggregation

// Last aggregations using generics for numeric types

type NumericLast[T any] struct {
	Base
	FieldName  string `json:"fieldName,omitempty"`
	TimeColumn string `json:"timeColumn,omitempty"`
}

func (n *NumericLast[T]) SetName(name string) *NumericLast[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericLast[T]) SetFieldName(fieldName string) *NumericLast[T] {
	n.FieldName = fieldName
	return n
}

func (n *NumericLast[T]) SetTimeColumn(timeColumn string) *NumericLast[T] {
	n.TimeColumn = timeColumn
	return n
}

// Type aliases for backward compatibility
type (
	DoubleLast = NumericLast[float64]
	FloatLast  = NumericLast[float32]
	LongLast   = NumericLast[int64]
)

// Constructor functions
func NewDoubleLast() *DoubleLast {
	d := &DoubleLast{}
	d.SetType("doubleLast")
	return d
}

func NewFloatLast() *FloatLast {
	f := &FloatLast{}
	f.SetType("floatLast")
	return f
}

func NewLongLast() *LongLast {
	l := &LongLast{}
	l.SetType("longLast")
	return l
}
