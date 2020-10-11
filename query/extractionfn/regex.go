package extractionfn

type Regex struct {
	Base
	Expr                    string `json:"expr"`
	Index                   int64  `json:"index"`
	ReplaceMissingValue     bool   `json:"replaceMissingValue"`
	ReplaceMissingValueWith string `json:"replaceMissingValueWith"`
}

func NewRegex() *Regex {
	r := &Regex{}
	r.SetType("regex")
	return r
}

func (r *Regex) SetExpr(expr string) *Regex {
	r.Expr = expr
	return r
}

func (r *Regex) SetIndex(index int64) *Regex {
	r.Index = index
	return r
}

func (r *Regex) SetReplaceMissingValue(replaceMissingValue bool) *Regex {
	r.ReplaceMissingValue = replaceMissingValue
	return r
}

func (r *Regex) SetReplaceMissingValueWith(replaceMissingValueWith string) *Regex {
	r.ReplaceMissingValueWith = replaceMissingValueWith
	return r
}
