package extractionfn

type Javascript struct {
	*Base
	Function  string `json:"function"`
	Injective bool   `json:"injective"`
}

func NewJavascript() *Javascript {
	j := &Javascript{}
	j.SetType("javascript")
	return j
}

func (j *Javascript) SetFunction(function string) *Javascript {
	j.Function = function
	return j
}

func (j *Javascript) SetInjective(injective bool) *Javascript {
	j.Injective = injective
	return j
}
