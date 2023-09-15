package query

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/toinclude"
)

type AnalysisType string

const (
	Cardinality      AnalysisType = "CARDINALITY"
	Size                          = "SIZE"
	Interval                      = "INTERVAL"
	Aggregators                   = "AGGREGATORS"
	MinMax                        = "MINMAX"
	TimestampSpec                 = "TIMESTAMPSPEC"
	QueryGranularity              = "QUERYGRANULARITY"
	Rollup                        = "ROLLUP"
)

type SegmentMetadata struct {
	Base
	ToInclude              builder.ToInclude `json:"toInclude,omitempty"`
	Merge                  *bool             `json:"merge,omitempty"`
	AnalysisTypes          []AnalysisType    `json:"analysisTypes,omitempty"`
	UsingDefaultInterval   *bool             `json:"usingDefaultInterval,omitempty"`
	LenientAggregatorMerge *bool             `json:"lenientAggregatorMerge,omitempty"`
}

func NewSegmentMetadata() *SegmentMetadata {
	s := &SegmentMetadata{}
	s.SetQueryType("segmentMetadata")
	return s
}

func (s *SegmentMetadata) SetDataSource(dataSource builder.DataSource) *SegmentMetadata {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *SegmentMetadata) SetIntervals(intervals builder.Intervals) *SegmentMetadata {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *SegmentMetadata) SetContext(context map[string]interface{}) *SegmentMetadata {
	s.Base.SetContext(context)
	return s
}

func (s *SegmentMetadata) SetToInclude(toInclude builder.ToInclude) *SegmentMetadata {
	s.ToInclude = toInclude
	return s
}

func (s *SegmentMetadata) SetMerge(merge bool) *SegmentMetadata {
	s.Merge = &merge
	return s
}

func (s *SegmentMetadata) SetAnalysisTypes(analysisTypes []AnalysisType) *SegmentMetadata {
	s.AnalysisTypes = analysisTypes
	return s
}

func (s *SegmentMetadata) SetUsingDefaultInterval(usingDefaultInterval bool) *SegmentMetadata {
	s.UsingDefaultInterval = &usingDefaultInterval
	return s
}

func (s *SegmentMetadata) SetLenientAggregatorMerge(lenientAggregatorMerge bool) *SegmentMetadata {
	s.LenientAggregatorMerge = &lenientAggregatorMerge
	return s
}

func (s *SegmentMetadata) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		ToInclude              json.RawMessage `json:"toInclude,omitempty"`
		Merge                  *bool           `json:"merge,omitempty"`
		AnalysisTypes          []AnalysisType  `json:"analysisTypes,omitempty"`
		UsingDefaultInterval   *bool           `json:"usingDefaultInterval,omitempty"`
		LenientAggregatorMerge *bool           `json:"lenientAggregatorMerge,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var t builder.ToInclude
	if tmp.ToInclude != nil {
		t, err = toinclude.Load(tmp.ToInclude)
		if err != nil {
			return err
		}
	}
	err = s.Base.UnmarshalJSON(data)
	s.ToInclude = t
	s.Merge = tmp.Merge
	s.AnalysisTypes = tmp.AnalysisTypes
	s.UsingDefaultInterval = tmp.UsingDefaultInterval
	s.LenientAggregatorMerge = tmp.LenientAggregatorMerge
	return err
}
