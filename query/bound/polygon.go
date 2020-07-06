package bound

type Polygon struct {
	*Base
	Abscissa []float64 `json:"abscissa"`
	Ordinate []float64 `json:"ordinate"`
}

func NewPolygon() *Polygon {
	p := &Polygon{}
	return p
}

func (p *Polygon) SetMinCoords(minCoords []float64) *Polygon {
	p.Base.SetMinCoords(minCoords)
	return p
}

func (p *Polygon) SetMaxCoords(maxCoords []float64) *Polygon {
	p.Base.SetMaxCoords(maxCoords)
	return p
}

func (p *Polygon) SetLimit(limit int64) *Polygon {
	p.Base.SetLimit(limit)
	return p
}

func (p *Polygon) SetAbscissa(abscissa []float64) *Polygon {
	p.Abscissa = abscissa
	return p
}

func (p *Polygon) SetOrdinate(ordinate []float64) *Polygon {
	p.Ordinate = ordinate
	return p
}
