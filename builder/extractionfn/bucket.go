package extractionfn

type Bucket struct {
	Base
	Size   *float64 `json:"size,omitempty"`
	Offset *float64 `json:"offset,omitempty"`
}

func NewBucket() *Bucket {
	b := &Bucket{}
	b.SetType("bucket")
	return b
}

func (b *Bucket) SetSize(size float64) *Bucket {
	b.Size = &size
	return b
}

func (b *Bucket) SetOffset(offset float64) *Bucket {
	b.Offset = &offset
	return b
}
