package datasource

type Union struct {
	Base
	DataSources []string `json:"dataSources,omitempty"`
}

func NewUnion() *Union {
	u := &Union{}
	u.SetType("union")
	return u
}

func (u *Union) SetDataSources(dataSources []string) *Union {
	u.DataSources = dataSources
	return u
}
