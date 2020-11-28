package aggregation

type StringLastFolding struct {
	Base
	FieldName      string `json:"fieldName,omitempty"`
	MaxStringBytes int64  `json:"maxStringBytes,omitempty"`
}

func NewStringLastFolding() *StringLastFolding {
	s := &StringLastFolding{}
	s.SetType("stringLastFolding")
	return s
}

func (s *StringLastFolding) SetName(name string) *StringLastFolding {
	s.Base.SetName(name)
	return s
}

func (s *StringLastFolding) SetFieldName(fieldName string) *StringLastFolding {
	s.FieldName = fieldName
	return s
}

func (s *StringLastFolding) SetMaxStringBytes(maxStringBytes int64) *StringLastFolding {
	s.MaxStringBytes = maxStringBytes
	return s
}
