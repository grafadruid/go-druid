package postaggregation

type HyperUniqueFinalizing struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
}

func NewHyperUniqueFinalizing() *HyperUniqueFinalizing {
	h := &HyperUniqueFinalizing{}
	h.SetType("hyperUniqueFinalizing")
	return h
}

func (h *HyperUniqueFinalizing) SetName(name string) *HyperUniqueFinalizing {
	h.Base.SetName(name)
	return h
}

func (h *HyperUniqueFinalizing) SetFieldName(fieldName string) *HyperUniqueFinalizing {
	h.FieldName = fieldName
	return h
}
