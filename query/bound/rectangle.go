package bound

type Rectangle struct {
	MinCoords []float64 `json:"minCoords"`
	MaxCoords []float64 `json:"maxCoords"`
	Limit     int64     `json:"limit"`
}

func NewRectangle() *Rectangle {
	r := &Rectangle{}
	return r
}

func (r *Rectangle) SetMinCoords(minCoords []float64) *Rectangle {
	r.MinCoords = minCoords
	return r
}

func (r *Rectangle) SetMaxCoords(maxCoords []float64) *Rectangle {
	r.MaxCoords = maxCoords
	return r
}

func (r *Rectangle) SetLimit(limit int64) *Rectangle {
	r.Limit = limit
	return r
}
