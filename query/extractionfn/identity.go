package extractionfn

type Identity struct {
	Base
}

func NewIdentity() *Identity {
	i := &Identity{}
	i.SetType("identity")
	return i
}
