package filter

type FilterTuning struct {
	Base
	UseBitmapIndex                 *bool `json:"useBitmapIndex,omitempty"`
	MinCardinalityToUseBitmapIndex int64 `json:"minCardinalityToUseBitmapIndex,omitempty"`
	MaxCardinalityToUseBitmapIndex int64 `json:"maxCardinalityToUseBitmapIndex,omitempty"`
}

func NewFilterTuning() *FilterTuning {
	f := &FilterTuning{}
	f.SetType("filterTuning")
	return f
}

func (f *FilterTuning) SetUseBitmapIndex(useBitmapIndex bool) *FilterTuning {
	f.UseBitmapIndex = &useBitmapIndex
	return f
}

func (f *FilterTuning) SetMinCardinalityToUseBitmapIndex(minCardinalityToUseBitmapIndex int64) *FilterTuning {
	f.MinCardinalityToUseBitmapIndex = minCardinalityToUseBitmapIndex
	return f
}

func (f *FilterTuning) SetMaxCardinalityToUseBitmapIndex(maxCardinalityToUseBitmapIndex int64) *FilterTuning {
	f.MaxCardinalityToUseBitmapIndex = maxCardinalityToUseBitmapIndex
	return f
}
