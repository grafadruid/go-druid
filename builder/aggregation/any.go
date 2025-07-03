package aggregation

// Any aggregations using generics for numeric types

type NumericAny[T any] struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func (n *NumericAny[T]) SetName(name string) *NumericAny[T] {
	n.Base.SetName(name)
	return n
}

func (n *NumericAny[T]) SetFieldName(fieldName string) *NumericAny[T] {
	n.FieldName = fieldName
	return n
}

// Type aliases for backward compatibility
type (
	DoubleAny = NumericAny[float64]
	FloatAny  = NumericAny[float32]
	LongAny   = NumericAny[int64]
)

// Constructor functions
func NewDoubleAny() *DoubleAny {
	d := &DoubleAny{}
	d.SetType("doubleAny")
	return d
}

func NewFloatAny() *FloatAny {
	f := &FloatAny{}
	f.SetType("floatAny")
	return f
}

func NewLongAny() *LongAny {
	l := &LongAny{}
	l.SetType("longAny")
	return l
}
