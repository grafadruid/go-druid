package types

type NullHandling string

const (
	NullString  NullHandling = "NULLSTRING"
	EmptyString              = "EMPTYSTRING"
	ReturnNull               = "RETURNNULL"
)
