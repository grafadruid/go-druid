package virtualcolumn

type OutputType string

const (
	Long   OutputType = "LONG"
	Float             = "FLOAT"
	Double            = "DOUBLE"
	String            = "STRING"
)

type Expression struct {
	Base
	Name       string `json:"name,omitempty"`
	Expression string `json:"expression,omitempty"`
	OutputType OutputType
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

func (e *Expression) SetOutputType(outputType OutputType) *Expression {
	e.OutputType = outputType
	return e
}
