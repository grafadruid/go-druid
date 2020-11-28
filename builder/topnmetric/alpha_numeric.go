package topnmetric

type AlphaNumeric struct {
	Base
	PreviousStop string `json:"previousStop,omitempty"`
}

func NewAlphaNumeric() *AlphaNumeric {
	a := &AlphaNumeric{}
	a.SetType("alphaNumeric")
	return a
}

func (a *AlphaNumeric) SetPreviousStop(previousStop string) *AlphaNumeric {
	a.PreviousStop = previousStop
	return a
}
