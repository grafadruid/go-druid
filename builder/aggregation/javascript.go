package aggregation

type Javascript struct {
	Base
	FieldNames  []string `json:"fieldNames,omitempty"`
	FnAggregate string   `json:"fnAggregate,omitempty"`
	FnReset     string   `json:"fnReset,omitempty"`
	FnCombine   string   `json:"fnCombine,omitempty"`
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

func (j *Javascript) SetFnAggregate(fnAggregate string) *Javascript {
	j.FnAggregate = fnAggregate
	return j
}

func (j *Javascript) SetFnReset(fnReset string) *Javascript {
	j.FnReset = fnReset
	return j
}

func (j *Javascript) SetFnCombine(fnCombine string) *Javascript {
	j.FnCombine = fnCombine
	return j
}
