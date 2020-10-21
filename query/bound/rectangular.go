package bound

type Rectangular struct {
	Base
	MinCoords []float64 `json:"minCoords"`
	MaxCoords []float64 `json:"maxCoords"`
	Limit     int64     `json:"limit"`
}

func NewRectangular() *Rectangular {
	r := &Rectangular{}
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
