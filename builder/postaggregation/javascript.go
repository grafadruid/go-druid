package postaggregation

type Javascript struct {
	Base
	FieldNames []string `json:"fieldNames,omitempty"`
	Function   string   `json:"function,omitempty"`
}

func NewJavascript() *Javascript {
	j := &Javascript{}
	j.SetType("javascript")
	return j
}

func (j *Javascript) SetName(name string) *Javascript {
	j.Base.SetName(name)
	return j
}

func (j *Javascript) SetFieldNames(fieldNames []string) *Javascript {
	j.FieldNames = fieldNames
	return j
}

func (j *Javascript) SetFunction(function string) *Javascript {
	j.Function = function
	return j
}
