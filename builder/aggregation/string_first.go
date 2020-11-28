package aggregation

type StringFirst struct {
	Base
	FieldName      string `json:"fieldName,omitempty"`
	MaxStringBytes int64  `json:"maxStringBytes,omitempty"`
}

func NewStringFirst() *StringFirst {
	s := &StringFirst{}
	s.SetType("stringFirst")
	return s
}

func (s *StringFirst) SetName(name string) *StringFirst {
	s.Base.SetName(name)
	return s
}

func (s *StringFirst) SetFieldName(fieldName string) *StringFirst {
	s.FieldName = fieldName
	return s
}

func (s *StringFirst) SetMaxStringBytes(maxStringBytes int64) *StringFirst {
	s.MaxStringBytes = maxStringBytes
	return s
}
