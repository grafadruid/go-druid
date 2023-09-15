package virtualcolumn

import "github.com/h2oai/go-druid/builder/types"

type Expression struct {
	Base
	Name       string           `json:"name,omitempty"`
	Expression string           `json:"expression,omitempty"`
	OutputType types.OutputType `json:"outputType,omitempty"`
}

func NewExpression() *Expression {
	e := &Expression{}
	e.SetType("expression")
	return e
}

func (e *Expression) SetName(name string) *Expression {
	e.Name = name
	return e
}

func (e *Expression) SetExpression(expression string) *Expression {
	e.Expression = expression
	return e
}

func (e *Expression) SetOutputType(outputType types.OutputType) *Expression {
	e.OutputType = outputType
	return e
}
