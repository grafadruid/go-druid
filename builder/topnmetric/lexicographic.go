package topnmetric

type Lexicographic struct {
	Base
	PreviousStop string `json:"previousStop,omitempty"`
}

func NewLexicographic() *Lexicographic {
	l := &Lexicographic{}
	l.SetType("lexicographic")
	return l
}

func (l *Lexicographic) SetPreviousStop(previousStop string) *Lexicographic {
	l.PreviousStop = previousStop
	return l
}
