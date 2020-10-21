package types

import "github.com/grafadruid/go-druid/query"

type mapLookupExtractor struct {
	Typ        string                 `json:"type"`
	Map        map[string]interface{} `json:"map"`
	IsOneToOne bool                   `json:"isOneToOne"`
}

func (m *mapLookupExtractor) Type() query.ComponentType {
	return m.Typ
}

// NewMapLookupExtractor instantiate a new map lookup extractor
func NewMapLookupExtractor(typ string, mapping map[string]interface{}, isOneToOne bool) query.LookupExtractor {
	return &mapLookupExtractor{Typ: typ, Map: mapping, IsOneToOne: isOneToOne}
}
