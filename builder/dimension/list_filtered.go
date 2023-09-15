package dimension

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/types"
)

type ListFiltered struct {
	Base
	Delegate    builder.Dimension `json:"delegate,omitempty"`
	Values      []string          `json:"values,omitempty"`
	IsWhiteList *bool             `json:"isWhiteList,omitempty"`
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

func (l *ListFiltered) SetDelegate(delegate builder.Dimension) *ListFiltered {
	l.Delegate = delegate
	return l
}

func (l *ListFiltered) SetValues(values []string) *ListFiltered {
	l.Values = values
	return l
}

func (l *ListFiltered) SetIsWhiteList(isWhiteList bool) *ListFiltered {
	l.IsWhiteList = &isWhiteList
	return l
}

func (l *ListFiltered) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Delegate    json.RawMessage `json:"delegate,omitempty"`
		Values      []string        `json:"values,omitempty"`
		IsWhiteList *bool           `json:"isWhiteList,omitempty"`
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
