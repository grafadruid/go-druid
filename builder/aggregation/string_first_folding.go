package aggregation

type StringFirstFolding struct {
	Base
	FieldName      string `json:"fieldName,omitempty"`
	MaxStringBytes int64  `json:"maxStringBytes,omitempty"`
}

func NewStringFirstFolding() *StringFirstFolding {
	s := &StringFirstFolding{}
	s.SetType("stringFirstFolding")
	return s
}

func (s *StringFirstFolding) SetName(name string) *StringFirstFolding {
	s.Base.SetName(name)
	return s
}

func (s *StringFirstFolding) SetFieldName(fieldName string) *StringFirstFolding {
	s.FieldName = fieldName
	return s
}

func (s *StringFirstFolding) SetMaxStringBytes(maxStringBytes int64) *StringFirstFolding {
	s.MaxStringBytes = maxStringBytes
	return s
}
