package filter

import (
	"encoding/json"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/extractionfn"
)

type Like struct {
	Base
	Dimension    string               `json:"dimension"`
	Pattern      string               `json:"pattern"`
	Escape       string               `json:"escapte"`
	ExtractionFn builder.ExtractionFn `json:"extractionFn"`
	FilterTuning *FilterTuning        `json:"filterTuning"`
}

func NewLike() *Like {
	l := &Like{}
	l.SetType("like")
	return l
}

func (l *Like) SetDimension(dimension string) *Like {
	l.Dimension = dimension
	return l
}

func (l *Like) SetPattern(pattern string) *Like {
	l.Pattern = pattern
	return l
}

func (l *Like) SetEscape(escape string) *Like {
	l.Escape = escape
	return l
}

func (l *Like) SetExtractionFn(extractionFn builder.ExtractionFn) *Like {
	l.ExtractionFn = extractionFn
	return l
}

func (l *Like) SetFilterTuning(filterTuning *FilterTuning) *Like {
	l.FilterTuning = filterTuning
	return l
}

func (l *Like) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Base
		Dimension    string          `json:"dimension"`
		Pattern      string          `json:"pattern"`
		Escape       string          `json:"escape"`
		ExtractionFn json.RawMessage `json:"extractionFn"`
		FilterTuning *FilterTuning   `json:"filterTuning"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	e, err := extractionfn.Load(tmp.ExtractionFn)
	if err != nil {
		return err
	}
	l.Base = tmp.Base
	l.Dimension = tmp.Dimension
	l.Pattern = tmp.Pattern
	l.Escape = tmp.Escape
	l.ExtractionFn = e
	l.FilterTuning = tmp.FilterTuning
	return nil
}
