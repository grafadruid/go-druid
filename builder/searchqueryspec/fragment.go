package searchqueryspec

type Fragment struct {
	Base
	Value         string `json:"value"`
	CaseSensitive bool   `json:"caseSensitive"`
}

func NewFragment() *Fragment {
	f := &Fragment{}
	f.SetType("fragment")
	return f
}

func (f *Fragment) SetValue(value string) *Fragment {
	f.Value = value
	return f
}

func (f *Fragment) SetCaseSensitive(caseSensitive bool) *Fragment {
	f.CaseSensitive = caseSensitive
	return f
}
