package extractionfn

type Lower struct {
	Base
	Locale string `json:"locale,omitempty"`
}

func NewLower() *Lower {
	l := &Lower{}
	l.SetType("lower")
	return l
}

func (l *Lower) SetLocale(locale string) *Lower {
	l.Locale = locale
	return l
}
