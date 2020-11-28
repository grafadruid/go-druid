package searchqueryspec

type InsensitiveContains struct {
	Base
	Value string `json:"value,omitempty"`
}

func NewInsensitiveContains() *InsensitiveContains {
	i := &InsensitiveContains{}
	i.SetType("insensitiveContains")
	return i
}

func (i *InsensitiveContains) SetValue(value string) *InsensitiveContains {
	i.Value = value
	return i
}
