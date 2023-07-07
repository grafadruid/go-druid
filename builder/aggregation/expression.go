package aggregation

// Expression holds the expression type of aggregator.
// TODO: Implement the backend changes.
type Expression struct {
	Base
	FieldsName            []string `json:"fields,omitempty"`
	AccumulatorIdentifier string   `json:"accumulatorIdentifier,omitempty"`
	InitialValue          string   `json:"initialValue,omitempty"`
	InitialCombineValue   string   `json:"initialCombineValue,omitempty"`
	Fold                  string   `json:"fold,omitempty"`
	Combine               string   `json:"combine,omitempty"`
	MaxSizeBytes          int64    `json:"maxSizeBytes,omitempty"`
}

// NewExpression create a new instance of Expression
func NewExpression() *Expression {
	t := &Expression{}
	t.Base.SetType("expression")
	return t
}

// SetName set name
func (t *Expression) SetName(name string) *Expression {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fields
func (t *Expression) SetFieldsName(fieldsName []string) *Expression {
	t.FieldsName = fieldsName
	return t
}

// SetAccumulatorIdentifier set accumulatorIdentifier
func (t *Expression) SetAccumulatorIdentifier(identifier string) *Expression {
	t.AccumulatorIdentifier = identifier
	return t
}

// SetInitialValue set initialValue
func (t *Expression) SetInitialValue(value string) *Expression {
	t.InitialValue = value
	return t
}

// SetInitialCombineValue set initialCombineValue
func (t *Expression) SetInitialCombineValue(value string) *Expression {
	t.InitialCombineValue = value
	return t
}

// SetFold set fold
func (t *Expression) SetFold(value string) *Expression {
	t.Fold = value
	return t
}

// SetCombine set combine
func (t *Expression) SetCombine(value string) *Expression {
	t.Combine = value
	return t
}

// SetMaxSizeBytes set caxSizeBytes
func (t *Expression) SetMaxSizeBytes(size int64) *Expression {
	t.MaxSizeBytes = size
	return t
}
