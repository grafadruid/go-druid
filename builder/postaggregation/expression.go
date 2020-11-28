package postaggregation

type Expression struct {
	Base
	Expression string `json:"expression,omitempty"`
	Ordering   string `json:"ordering,omitempty"`
}

func NewExpression() *Expression {
	e := &Expression{}
	e.SetType("expression")
	return e
}

func (e *Expression) SetName(name string) *Expression {
	e.Base.SetName(name)
	return e
}

func (e *Expression) SetExpression(expression string) *Expression {
	e.Expression = expression
	return e
}

func (e *Expression) SetOrdering(ordering string) *Expression {
	e.Ordering = ordering
	return e
}
