package extractionfn

type Partial struct {
	Base
	Expr string `json:"expr,omitempty"`
}

func NewPartial() *Partial {
	p := &Partial{}
	p.SetType("partial")
	return p
}

func (p *Partial) SetExpr(expr string) *Partial {
	p.Expr = expr
	return p
}
