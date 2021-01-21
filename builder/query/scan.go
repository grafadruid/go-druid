package query

import (
	"encoding/json"
	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/filter"
	"github.com/grafadruid/go-druid/builder/virtualcolumn"
)

type Order string

const (
	Ascending  Order = "ASCENDING"
	Descending       = "DESCENDING"
	None             = "NONE"
)

type Scan struct {
	Base
	VirtualColumns []builder.VirtualColumn `json:"virtualColumns,omitempty"`
	ResultFormat   string                  `json:"resultFormat,omitempty"`
	BatchSize      int64                   `json:"batchSize,omitempty"`
	Limit          int64                   `json:"limit,omitempty"`
	Order          Order                   `json:"order,omitempty"`
	Filter         builder.Filter          `json:"filter,omitempty"`
	Columns        []string                `json:"columns,omitempty"`
	Legacy         bool                    `json:"legacy,omitempty"`
}

func NewScan() *Scan {
	s := &Scan{}
	s.Base.SetQueryType("scan")
	return s
}

func (s *Scan) SetDataSource(dataSource builder.DataSource) *Scan {
	s.Base.SetDataSource(dataSource)
	return s
}

func (s *Scan) SetIntervals(intervals builder.Intervals) *Scan {
	s.Base.SetIntervals(intervals)
	return s
}

func (s *Scan) SetContext(context map[string]interface{}) *Scan {
	s.Base.SetContext(context)
	return s
}

func (s *Scan) SetVirtualColumns(virtualColumns []builder.VirtualColumn) *Scan {
	s.VirtualColumns = virtualColumns
	return s
}

func (s *Scan) SetResultFormat(resultFormat string) *Scan {
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

func (s *Scan) SetFilter(filter builder.Filter) *Scan {
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

func (s *Scan) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		VirtualColumns []json.RawMessage `json:"virtualColumns,omitempty"`
		ResultFormat   string            `json:"resultFormat,omitempty"`
		BatchSize      int64             `json:"batchSize,omitempty"`
		Limit          int64             `json:"limit,omitempty"`
		Order          Order             `json:"order,omitempty"`
		Filter         json.RawMessage   `json:"filter,omitempty"`
		Columns        []string          `json:"columns,omitempty"`
		Legacy         bool              `json:"legacy,omitempty"`
	}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	var v builder.VirtualColumn
	vv := make([]builder.VirtualColumn, len(tmp.VirtualColumns))
	for i := range tmp.VirtualColumns {
		if v, err = virtualcolumn.Load(tmp.VirtualColumns[i]); err != nil {
			return err
		}
		vv[i] = v
	}
	var f builder.Filter
	if tmp.Filter != nil {
		f, err = filter.Load(tmp.Filter)
		if err != nil {
			return err
		}
	}
	err = s.Base.UnmarshalJSON(data)
	s.VirtualColumns = vv
	s.ResultFormat = tmp.ResultFormat
	s.BatchSize = tmp.BatchSize
	s.Limit = tmp.Limit
	s.Order = tmp.Order
	s.Filter = f
	s.Columns = tmp.Columns
	s.Legacy = tmp.Legacy
	return err
}
