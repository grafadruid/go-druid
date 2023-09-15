package query

import (
	"encoding/json"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/filter"
	"github.com/h2oai/go-druid/builder/virtualcolumn"
)

type Order string

const (
	Ascending  Order = "ASCENDING"
	Descending       = "DESCENDING"
	None             = "NONE"
)

// Scan query returns raw Apache Druid rows in streaming mode.
// https://druid.apache.org/docs/latest/querying/scan-query.html
type Scan struct {
	Base
	VirtualColumns []builder.VirtualColumn `json:"virtualColumns,omitempty"`
	ResultFormat   string                  `json:"resultFormat,omitempty"`
	BatchSize      int64                   `json:"batchSize,omitempty"`
	Limit          int64                   `json:"limit,omitempty"`
	Offset         int64                   `json:"offset,omitempty"`
	Order          Order                   `json:"order,omitempty"`
	Filter         builder.Filter          `json:"filter,omitempty"`
	Columns        []string                `json:"columns,omitempty"`
	Legacy         *bool                   `json:"legacy,omitempty"`
}

// NewScan returns *builder.Scan which can be used to build a scan query.
// Eg,
//
//	table := datasource.NewTable().SetName("table-name")
//
//	now := time.Now()
//	i := intervals.NewInterval().SetInterval(now.Add(-60*time.Minute), now)
//	is := intervals.NewIntervals().SetIntervals([]*intervals.Interval{i})
//
//	filter1 := filter.NewSelector().SetDimension("key1").SetValue("val1")
//	filter2 := filter.NewSelector().SetDimension("key2").SetValue("val2")
//	filters := filter.NewAnd().SetFields([]builder.Filter{filter1, filter2})
//
//	ts := query.NewScan().SetDataSource(table).SetIntervals(is).SetFilter(filters).SetResultFormat("compactedList").SetLimit(10)
func NewScan() *Scan {
	s := &Scan{}
	s.Base.SetQueryType("scan")
	return s
}

// SetDataSource sets data source.
func (s *Scan) SetDataSource(dataSource builder.DataSource) *Scan {
	s.Base.SetDataSource(dataSource)
	return s
}

// SetIntervals set the intervals.
func (s *Scan) SetIntervals(intervals builder.Intervals) *Scan {
	s.Base.SetIntervals(intervals)
	return s
}

// SetContext sets the context.
func (s *Scan) SetContext(context map[string]interface{}) *Scan {
	s.Base.SetContext(context)
	return s
}

// SetVirtualColumns sets virtual columns.
func (s *Scan) SetVirtualColumns(virtualColumns []builder.VirtualColumn) *Scan {
	s.VirtualColumns = virtualColumns
	return s
}

// SetResultFormat sets the result format.
func (s *Scan) SetResultFormat(resultFormat string) *Scan {
	s.ResultFormat = resultFormat
	return s
}

// SetBatchSize sets the batch size.
func (s *Scan) SetBatchSize(batchSize int64) *Scan {
	s.BatchSize = batchSize
	return s
}

// SetLimit sets the limit.
func (s *Scan) SetLimit(limit int64) *Scan {
	s.Limit = limit
	return s
}

// SetOffset sets the offset.
func (s *Scan) SetOffset(offset int64) *Scan {
	s.Offset = offset
	return s
}

// SetOrder sets the order.
func (s *Scan) SetOrder(order Order) *Scan {
	s.Order = order
	return s
}

// SetFilter sets the filter.
func (s *Scan) SetFilter(filter builder.Filter) *Scan {
	s.Filter = filter
	return s
}

// SetColumns set columns.
func (s *Scan) SetColumns(columns []string) *Scan {
	s.Columns = columns
	return s
}

// SetLegacy sets the `druid.query.scan.legacy` field.
func (s *Scan) SetLegacy(legacy bool) *Scan {
	s.Legacy = &legacy
	return s
}

// UnmarshalJSON unmarshalls a druid scan native query json string into builder type.
func (s *Scan) UnmarshalJSON(data []byte) error {
	var err error
	var tmp struct {
		VirtualColumns []json.RawMessage `json:"virtualColumns,omitempty"`
		ResultFormat   string            `json:"resultFormat,omitempty"`
		BatchSize      int64             `json:"batchSize,omitempty"`
		Limit          int64             `json:"limit,omitempty"`
		Offset         int64             `json:"offset,omitempty"`
		Order          Order             `json:"order,omitempty"`
		Filter         json.RawMessage   `json:"filter,omitempty"`
		Columns        []string          `json:"columns,omitempty"`
		Legacy         *bool             `json:"legacy,omitempty"`
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
	s.Offset = tmp.Offset
	s.Order = tmp.Order
	s.Filter = f
	s.Columns = tmp.Columns
	s.Legacy = tmp.Legacy
	return err
}
