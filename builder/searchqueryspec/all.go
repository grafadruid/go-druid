package searchqueryspec

type All struct {
	Base
}

func NewAll() *All {
	a := &All{}
	a.SetType("all")
	return a
}
