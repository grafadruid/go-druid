package aggregation

type StringLast struct {
	*Base
	FieldName      string `json:"fieldName"`
	MaxStringBytes int64  `json:"maxStringBytes"`
}

func NewStringLast() *StringLast {
	s := &StringLast{}
	s.SetType("stringLast")
	return s
}

func (s *StringLast) SetName(name string) *StringLast {
	s.Base.SetName(name)
	return s
}

func (s *StringLast) SetFieldName(fieldName string) *StringLast {
	s.FieldName = fieldName
	return s
}

func (s *StringLast) SetMaxStringBytes(maxStringBytes int64) *StringLast {
	s.MaxStringBytes = maxStringBytes
	return s
}
