package topnmetric

type Numeric struct {
	Base
	Metric string `json:"metric,omitempty"`
}

func NewNumeric() *Numeric {
	n := &Numeric{}
	n.SetType("numeric")
	return n
}

func (n *Numeric) SetMetric(metric string) *Numeric {
	n.Metric = metric
	return n
}
