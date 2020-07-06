package druid

const (
	StatusEndpoint         = "status"
	HealthEndpoint         = "status/health"
	PropertiesEndpoint     = "status/properties"
	SelfDiscoveredEndpoint = "status/selfDiscovered/status"
)

type Status struct {
	Version string `json:"version"`
	Modules []struct {
		Name     string `json:"name"`
		Artifact string `json:"artifact"`
		Version  string `json:"version"`
	} `json:"modules"`
	Memory struct {
		MaxMemory    int `json:"maxMemory"`
		TotalMemory  int `json:"totalMemory"`
		FreeMemory   int `json:"freeMemory"`
		UsedMemory   int `json:"usedMemory"`
		DirectMemory int `json:"directMemory"`
	} `json:"memory"`
}
type Health bool
type Properties map[string]string
type SelfDiscovered struct {
	SelfDiscovered bool `json:"selfDiscovered"`
}

type CommonService struct {
	client *Client
}

func (c *CommonService) Status() (*Status, *Response, error) {
	var s *Status
	response, err := c.client.ExecuteRequest("GET", StatusEndpoint, nil, &s)
	if err != nil {
		return nil, response, err
	}
	return s, response, nil
}

func (c *CommonService) Health() (*Health, *Response, error) {
	var h *Health
	response, err := c.client.ExecuteRequest("GET", HealthEndpoint, nil, &h)
	if err != nil {
		return nil, response, err
	}
	return h, response, nil
}

func (c *CommonService) Properties() (*Properties, *Response, error) {
	var p *Properties
	response, err := c.client.ExecuteRequest("GET", PropertiesEndpoint, nil, &p)
	if err != nil {
		return nil, response, err
	}
	return p, response, nil
}

func (c *CommonService) SelfDiscovered() (*SelfDiscovered, *Response, error) {
	var s *SelfDiscovered
	response, err := c.client.ExecuteRequest("GET", SelfDiscoveredEndpoint, nil, &s)
	if err != nil {
		return nil, response, err
	}
	return s, response, nil
}
