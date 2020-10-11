package filter

import "github.com/grafadruid/go-druid/query"

type Like struct {
	Base
	Dimension    string             `json:"dimension"`
	Pattern      string             `json:"pattern"`
	Escape       string             `json:"escapte,omitempty"`
	ExtractionFn query.ExtractionFn `json:"extractionFn,omitempty"`
	FilterTuning *FilterTuning      `json:"filterTuning,omitempty"`
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

func (l *Like) SetExtractionFn(extractionFn query.ExtractionFn) *Like {
	l.ExtractionFn = extractionFn
	return l
}

func (l *Like) SetFilterTuning(filterTuning *FilterTuning) *Like {
	l.FilterTuning = filterTuning
	return l
}
