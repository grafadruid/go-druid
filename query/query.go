package query

type QueryLanguage int
type ComponentType = string

const (
	NativeLanguage QueryLanguage = iota
	SQLLanguage
)

type Query interface {
	Type() ComponentType
	Language() QueryLanguage
}

type Aggregator interface {
	Type() ComponentType
}

type Bound interface {
	Type() ComponentType
}

type DataSource interface {
	Type() ComponentType
}

type Dimension interface {
	Type() ComponentType
}

type ExtractionFn interface {
	Type() ComponentType
}

type Filter interface {
	Type() ComponentType
}

type Granularity interface {
	Type() ComponentType
}

type HavingSpec interface {
	Type() ComponentType
}

type LimitSpec interface {
	Type() ComponentType
}

type LookupExtractor interface {
	Type() ComponentType
}

type PostAggregator interface {
	Type() ComponentType
}

type SearchQuerySpec interface {
	Type() ComponentType
}

type TopNMetric interface {
	Type() ComponentType
}

type VirtualColumn interface {
	Type() ComponentType
}
