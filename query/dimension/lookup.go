package dimension

import "github.com/grafadruid/go-druid/query"

type MapLookup struct {
	*Base
	ReplaceMissingValueWith string                `json:"replaceMissingValueWith"`
	RetainMissingValue      bool                  `json:"retainMissingValue"`
	Lookup                  query.LookupExtractor `json:"lookup"`
	Optimize                bool                  `json:"optimize"`
}

type RegisteredLookup struct {
	*Base
	Name string `json:"name"`
}
