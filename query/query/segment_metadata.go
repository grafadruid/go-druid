package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
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
	ToInclude              string         `json:"toInclude"`
	Merge                  bool           `json:"merge"`
	AnalysisTypes          []AnalysisType `json:"analysisTypes"`
	UsingDefaultInterval   bool           `json:"usingDefaultInterval"`
	LenientAggregatorMerge bool           `json:"lenientAggregatorMerge"`
}

func NewSegmentMetadata() *SegmentMetadata {
	s := &SegmentMetadata{}
	s.SetQueryType("segmentMetadata")
	return s
}

func (s *SegmentMetadata) SetDataSource(dataSource query.DataSource) *SegmentMetadata {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *SegmentMetadata) SetIntervals(intervals []types.Interval) *SegmentMetadata {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *SegmentMetadata) SetContext(context map[string]interface{}) *SegmentMetadata {
	s.Base.SetContext(context)
	return s
}

func (s *SegmentMetadata) SetToInclude(toInclude string) *SegmentMetadata {
	s.ToInclude = toInclude
	return s
}

func (s *SegmentMetadata) SetMerge(merge bool) *SegmentMetadata {
	s.Merge = merge
	return s
}

func (s *SegmentMetadata) SetAnalysisTypes(analysisTypes []AnalysisType) *SegmentMetadata {
	s.AnalysisTypes = analysisTypes
	return s
}

func (s *SegmentMetadata) SetUsingDefaultInterval(usingDefaultInterval bool) *SegmentMetadata {
	s.UsingDefaultInterval = usingDefaultInterval
	return s
}

func (s *SegmentMetadata) SetLenientAggregatorMerge(lenientAggregatorMerge bool) *SegmentMetadata {
	s.LenientAggregatorMerge = lenientAggregatorMerge
	return s
}
