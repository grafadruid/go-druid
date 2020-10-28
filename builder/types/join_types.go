package types

type JoinType string

const (
	Inner JoinType = "INNER"
	Left           = "LEFT"
	Right          = "RIGHT"
	Full           = "FULL"
)
