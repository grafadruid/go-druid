package aggregation

// String folding aggregations using generics

type StringFolding struct {
	Base
	FieldName      string `json:"fieldName,omitempty"`
	MaxStringBytes int64  `json:"maxStringBytes,omitempty"`
}

func (s *StringFolding) SetName(name string) *StringFolding {
	s.Base.SetName(name)
	return s
}

func (s *StringFolding) SetFieldName(fieldName string) *StringFolding {
	s.FieldName = fieldName
	return s
}

func (s *StringFolding) SetMaxStringBytes(maxStringBytes int64) *StringFolding {
	s.MaxStringBytes = maxStringBytes
	return s
}

// Type aliases for backward compatibility
type (
	StringFirstFolding = StringFolding
	StringLastFolding  = StringFolding
)

// Constructor functions
func NewStringFirstFolding() *StringFirstFolding {
	s := &StringFirstFolding{}
	s.SetType("stringFirstFolding")
	return s
}

func NewStringLastFolding() *StringLastFolding {
	s := &StringLastFolding{}
	s.SetType("stringLastFolding")
	return s
}
