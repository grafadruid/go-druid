package lookup

type Map struct {
	Base
	Map        map[string]string `json:"map"`
	IsOneToOne bool              `json:"isOneToOne"`
}

func NewMap() *Map {
	m := &Map{}
	m.SetType("map")
	return m
}

func (m *Map) SetMap(mp map[string]string) *Map {
	m.Map = mp
	return m
}

func (m *Map) SetIsOneToOne(b bool) *Map {
	m.IsOneToOne = b
	return m
}
