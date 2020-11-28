package bound

type Rectangular struct {
	Base
	MinCoords []float64 `json:"minCoords,omitempty"`
	MaxCoords []float64 `json:"maxCoords,omitempty"`
	Limit     int64     `json:"limit,omitempty"`
}

func NewRectangular() *Rectangular {
	r := &Rectangular{}
	r.SetType("rectangular")
	return r
}

func (r *Rectangular) SetMinCoords(minCoords []float64) *Rectangular {
	r.MinCoords = minCoords
	return r
}

func (r *Rectangular) SetMaxCoords(maxCoords []float64) *Rectangular {
	r.MaxCoords = maxCoords
	return r
}

func (r *Rectangular) SetLimit(limit int64) *Rectangular {
	r.Limit = limit
	return r
}
