package query

type QueryLanguage int

const (
	NativeLanguage QueryLanguage = iota
	SQLLanguage
)

type Query interface {
	ID() string
	Language() QueryLanguage
}

type Aggregator interface{}

type Bound interface{}

type DataSource interface{}

type Dimension interface{}

type ExtractionFn interface{}

type Filter interface{}

type Granularity interface{}

type HavingSpec interface{}

type LimitSpec interface{}

type LookupExtractor interface{}

type PostAggregator interface{}

type SearchQuerySpec interface{}

type TopNMetric interface{}

type VirtualColumn interface{}
