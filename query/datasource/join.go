package datasource

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Join struct {
	Base
	Left        query.DataSource `json:"left"`
	Right       query.DataSource `json:"right"`
	RightPrefix string           `json:"rightPrefix"`
	Condition   string           `json:"condition"`
	JoinType    types.JoinType   `json:"joinType"`
}

func NewJoin() *Join {
	j := &Join{}
	j.SetType("join")
	return j
}

func (j *Join) SetLeft(left query.DataSource) *Join {
	j.Left = left
	return j
}

func (j *Join) SetRight(right query.DataSource) *Join {
	j.Right = right
	return j
}

func (j *Join) SetRightPrefix(rightPrefix string) *Join {
	j.RightPrefix = rightPrefix
	return j
}

func (j *Join) SetCondition(condition string) *Join {
	j.Condition = condition
	return j
}

func (j *Join) SetJoinType(joinType types.JoinType) *Join {
	j.JoinType = joinType
	return j
}
