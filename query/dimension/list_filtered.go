package dimension

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type ListFiltered struct {
	Base
	Delegate    query.Dimension `json:"delegate"`
	Values      []string        `json:"values"`
	IsWhiteList bool            `json:"isWhiteList"`
}

func NewListFiltered() *ListFiltered {
	l := &ListFiltered{}
	l.SetType("listFiltered")
	return l
}

func (l *ListFiltered) SetDimension(dimension string) *ListFiltered {
	l.Base.SetDimension(dimension)
	return l
}

func (l *ListFiltered) SetOutputName(outputName string) *ListFiltered {
	l.Base.SetOutputName(outputName)
	return l
}

func (l *ListFiltered) SetOutputType(outputType types.OutputType) *ListFiltered {
	l.Base.SetOutputType(outputType)
	return l
}

func (l *ListFiltered) SetDelegate(delegate query.Dimension) *ListFiltered {
	l.Delegate = delegate
	return l
}

func (l *ListFiltered) SetValues(values []string) *ListFiltered {
	l.Values = values
	return l
}

func (l *ListFiltered) SetIsWhiteList(isWhiteList bool) *ListFiltered {
	l.IsWhiteList = isWhiteList
	return l
}

func (l *ListFiltered) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Delegate    json.RawMessage `json:"delegate"`
		Values      []string        `json:"values"`
		IsWhiteList bool            `json:"isWhiteList"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	d, err := Load(tmp.Delegate)
	if err != nil {
		return err
	}
	l.Base = tmp.Base
	l.Delegate = d
	l.Values = tmp.Values
	l.IsWhiteList = tmp.IsWhiteList
	return nil
}
