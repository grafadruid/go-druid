package common

import "github.com/grafadruid/go-druid/builder"

// TypeOnlyBase provides shared base struct for components that only need a type field.
// This eliminates code duplication across filter, extractionfn, havingspec, topnmetric,
// searchqueryspec, datasource, bound, limitspec, virtualcolumn, and intervals packages.
type TypeOnlyBase struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

// SetType sets the component type and returns the base for method chaining
func (b *TypeOnlyBase) SetType(typ builder.ComponentType) *TypeOnlyBase {
	b.Typ = typ
	return b
}

// Type returns the component type
func (b *TypeOnlyBase) Type() builder.ComponentType {
	return b.Typ
}

// NamedBase provides shared base struct for components that need both type and name fields.
// This eliminates code duplication across aggregation and postaggregation packages.
type NamedBase struct {
	Typ  builder.ComponentType `json:"type,omitempty"`
	Name string                `json:"name,omitempty"`
}

// SetType sets the component type and returns the base for method chaining
func (b *NamedBase) SetType(typ builder.ComponentType) *NamedBase {
	b.Typ = typ
	return b
}

// SetName sets the component name and returns the base for method chaining
func (b *NamedBase) SetName(name string) *NamedBase {
	b.Name = name
	return b
}

// Type returns the component type
func (b *NamedBase) Type() builder.ComponentType {
	return b.Typ
}
