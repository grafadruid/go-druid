package lookup

type Map struct {
	Base
	Map        map[string]string `json:"map,omitempty"`
	IsOneToOne *bool             `json:"isOneToOne,omitempty"`
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

func (m *Map) SetIsOneToOne(i bool) *Map {
	m.IsOneToOne = &i
	return m
}
