package aggregation

type StringAny struct {
	Base
	FieldName      string `json:"fieldName,omitempty"`
	MaxStringBytes int64  `json:"maxStringBytes,omitempty"`
}

func NewStringAny() *StringAny {
	s := &StringAny{}
	s.SetType("stringAny")
	return s
}

func (s *StringAny) SetName(name string) *StringAny {
	s.Base.SetName(name)
	return s
}

func (s *StringAny) SetFieldName(fieldName string) *StringAny {
	s.FieldName = fieldName
	return s
}

func (s *StringAny) SetMaxStringBytes(maxStringBytes int64) *StringAny {
	s.MaxStringBytes = maxStringBytes
	return s
}
