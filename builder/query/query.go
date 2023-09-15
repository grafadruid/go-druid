package query

import (
	"encoding/json"
	"errors"

	"github.com/h2oai/go-druid/builder"
	"github.com/h2oai/go-druid/builder/datasource"
	"github.com/h2oai/go-druid/builder/intervals"
)

type Base struct {
	ID         string                 `json:"ID,omitempty"`
	QueryType  builder.ComponentType  `json:"queryType,omitempty"`
	DataSource builder.DataSource     `json:"dataSource,omitempty"`
	Intervals  builder.Intervals      `json:"intervals,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

func (b *Base) SetID(ID string) *Base {
	b.ID = ID
	return b
}

func (b *Base) SetQueryType(queryType builder.ComponentType) *Base {
	b.QueryType = queryType
	return b
}

func (b *Base) SetDataSource(dataSource builder.DataSource) *Base {
	b.DataSource = dataSource
	return b
}

func (b *Base) SetIntervals(intervals builder.Intervals) *Base {
	b.Intervals = intervals
	return b
}

func (b *Base) SetContext(context map[string]interface{}) *Base {
	b.Context = context
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.QueryType
}

func (b *Base) UnmarshalJSON(data []byte) error {
	var tmp struct {
		ID         string                 `json:"ID,omitempty"`
		QueryType  builder.ComponentType  `json:"queryType,omitempty"`
		DataSource json.RawMessage        `json:"dataSource,omitempty"`
		Intervals  json.RawMessage        `json:"intervals,omitempty"`
		Context    map[string]interface{} `json:"context,omitempty"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if b.Type() != "sql" {
		d, err := datasource.Load(tmp.DataSource)
		if err != nil {
			return err
		}
		if d.Type() == "query" {
			d.(*datasource.Query).UnmarshalJSONWithQueryLoader(tmp.DataSource, Load)
		}
		b.DataSource = d
		var i builder.Intervals
		if tmp.Intervals != nil {
			i, err = intervals.Load(tmp.Intervals)
			if err != nil {
				return err
			}
		}
		b.Intervals = i
	}
	b.ID = tmp.ID
	b.QueryType = tmp.QueryType
	b.Context = tmp.Context
	return nil
}

func Load(data []byte) (builder.Query, error) {
	var q builder.Query
	if string(data) == "null" {
		return q, nil
	}
	var t struct {
		Typ builder.ComponentType `json:"queryType,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	switch t.Typ {
	case "dataSourceMetadata":
		q = NewDataSourceMetadata()
	case "groupBy":
		q = NewGroupBy()
	case "scan":
		q = NewScan()
	case "search":
		q = NewSearch()
	case "segmentMetadata":
		q = NewSegmentMetadata()
	case "sql":
		q = NewSQL()
	case "timeBoundary":
		q = NewTimeBoundary()
	case "timeseries":
		q = NewTimeseries()
	case "topN":
		q = NewTopN()
	default:
		return nil, errors.New("unsupported query type")
	}
	return q, json.Unmarshal(data, &q)
}
