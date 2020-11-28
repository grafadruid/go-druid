package extractionfn

type Substring struct {
	Base
	Index  int64 `json:"index,omitempty"`
	Length int64 `json:"length,omitempty"`
}

func NewSubstring() *Substring {
	s := &Substring{}
	s.SetType("substring")
	return s
}

func (s *Substring) SetIndex(index int64) *Substring {
	s.Index = index
	return s
}

func (s *Substring) SetLength(length int64) *Substring {
	s.Length = length
	return s
}
