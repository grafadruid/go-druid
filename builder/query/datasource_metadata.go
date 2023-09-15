package query

import (
	"github.com/h2oai/go-druid/builder"
)

type DataSourceMetadata struct {
	Base
}

func NewDataSourceMetadata() *DataSourceMetadata {
	d := &DataSourceMetadata{}
	d.SetQueryType("dataSourceMetadata")
	return d
}

func (d *DataSourceMetadata) SetDataSource(dataSource builder.DataSource) *DataSourceMetadata {
	d.Base.SetDataSource(dataSource)
	return d
}

func (d *DataSourceMetadata) SetIntervals(intervals builder.Intervals) *DataSourceMetadata {
	d.Base.SetIntervals(intervals)
	return d
}

func (d *DataSourceMetadata) SetContext(context map[string]interface{}) *DataSourceMetadata {
	d.Base.SetContext(context)
	return d
}
