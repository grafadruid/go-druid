package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type Order string

const (
	Ascending  Order = "ASCENDING"
	Descending       = "DESCENDING"
	None             = "NONE"
)

type Scan struct {
	*Base
	VirtualColumns []query.VirtualColumn `json:"virtualColumns"`
	ResultFormat   types.ResultFormat    `json:"resultFormat"`
	BatchSize      int64                 `json:"batchSize"`
	Limit          int64                 `json:"limit"`
	Order          Order                 `json:"order"`
	Filter         query.Filter          `json:"filter"`
	Columns        []string              `json:"columns"`
	Legacy         bool                  `json:"legacy"`
}

func NewScan() *Scan {
	s := &Scan{}
	s.SetQueryType("scan")
	return s
}

func (s *Scan) SetDataSource(dataSource query.DataSource) *Scan {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *Scan) SetIntervals(intervals []types.Interval) *Scan {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *Scan) SetContext(context map[string]interface{}) *Scan {
	s.Base.SetContext(context)
	return s
}

func (s *Scan) SetVirtualColumns(virtualColumns []query.VirtualColumn) *Scan {
	s.VirtualColumns = virtualColumns
	return s
}

func (s *Scan) SetResultFormat(resultFormat types.ResultFormat) *Scan {
	s.ResultFormat = resultFormat
	return s
}

func (s *Scan) SetBatchSize(batchSize int64) *Scan {
	s.BatchSize = batchSize
	return s
}

func (s *Scan) SetLimit(limit int64) *Scan {
	s.Limit = limit
	return s
}

func (s *Scan) SetOrder(order Order) *Scan {
	s.Order = order
	return s
}

func (s *Scan) SetFilter(filter query.Filter) *Scan {
	s.Filter = filter
	return s
}

func (s *Scan) SetColumns(columns []string) *Scan {
	s.Columns = columns
	return s
}

func (s *Scan) SetLegacy(legacy bool) *Scan {
	s.Legacy = legacy
	return s
}
