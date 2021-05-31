package searchqueryspec

type Contains struct {
	Base
	Value         string `json:"value,omitempty"`
	CaseSensitive *bool  `json:"caseSensitive,omitempty"`
}

func NewContains() *Contains {
	c := &Contains{}
	c.SetType("contains")
	return c
}

func (c *Contains) SetValue(value string) *Contains {
	c.Value = value
	return c
}

func (c *Contains) SetCaseSensitive(caseSensitive bool) *Contains {
	c.CaseSensitive = &caseSensitive
	return c
}
