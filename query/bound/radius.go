package bound

type Radius struct {
	Base
	Coords []float64 `json:"coords"`
	Radius float64   `json:"radius"`
}

func NewRadius() *Radius {
	r := &Radius{}
	return r
}

func (r *Radius) SetCoords(coords []float64) *Radius {
	r.Coords = coords
	return r
}

func (r *Radius) SetRadius(radius float64) *Radius {
	r.Radius = radius
	return r
}
