package filter

type Expression struct {
	Base
	Expression   string        `json:"expression,omitempty"`
	FilterTuning *FilterTuning `json:"filterTuning,omitempty"`
}

func NewExpression() *Expression {
	e := &Expression{}
	e.SetType("expression")
	return e
}

func (e *Expression) SetExpression(expression string) *Expression {
	e.Expression = expression
	return e
}

func (e *Expression) SetFilterTuning(filterTuning *FilterTuning) *Expression {
	e.FilterTuning = filterTuning
	return e
}
