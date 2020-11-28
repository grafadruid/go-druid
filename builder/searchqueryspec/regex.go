package searchqueryspec

type Regex struct {
	Base
	Pattern string `json:"pattern,omitempty"`
}

func NewRegex() *Regex {
	r := &Regex{}
	r.SetType("regex")
	return r
}

func (r *Regex) SetPattern(pattern string) *Regex {
	r.Pattern = pattern
	return r
}
