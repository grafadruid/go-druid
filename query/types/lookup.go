package types

import "github.com/grafadruid/go-druid/query"

type mapLookupExtractor struct {
	Type       string                 `json:"type"`
	Map        map[string]interface{} `json:"map"`
	IsOneToOne bool                   `json:"isOneToOne"`
}

// NewMapLookupExtractor instantiate a new map lookup extractor
func NewMapLookupExtractor(typ string, mapping map[string]interface{}, isOneToOne bool) query.LookupExtractor {
	return &mapLookupExtractor{Type: typ, Map: mapping, IsOneToOne: isOneToOne}
}
