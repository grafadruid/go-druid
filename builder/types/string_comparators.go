package types

type StringComparator string

const (
	Lexicographic StringComparator = "lexicographic"
	Alphanumeric                   = "alphanumeric"
	Numeric                        = "numeric"
	Strlen                         = "strlen"
	Version                        = "version"
)
