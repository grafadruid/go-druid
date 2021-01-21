package builder

type ComponentType = string

type Query interface {
	Type() ComponentType
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

type ToInclude interface {
	Type() ComponentType
}

type VirtualColumn interface {
	Type() ComponentType
}

type Intervals interface {
	Type() ComponentType
}
