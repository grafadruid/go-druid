package extractionfn

type Upper struct {
	Base
	Locale string `json:"locale,omitempty"`
}

func NewUpper() *Upper {
	u := &Upper{}
	u.SetType("upper")
	return u
}

func (u *Upper) SetLocale(locale string) *Upper {
	u.Locale = locale
	return u
}
