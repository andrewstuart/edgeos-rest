package edgeos

// PortForward is a struct that represents a port forwarding rule
type PortForward struct {
	PortFrom    string `json:"original-port"`
	PortTo      string `json:"forward-to-port"`
	IPTo        string `json:"forward-to-address"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}

// PortForwards is a struct that represents the response from the API for with
// the list of port forwards, etc.
type PortForwards struct {
	AutoFirewall string        `json:"auto-firewall"`
	HairpinNAT   string        `json:"hairpin-nat"`
	WAN          string        `json:"wan"`
	Lans         []LanConfig   `json:"lans-config"`
	Rules        []PortForward `json:"rules-config"`
}

// LanConfig is a simple string map that encapsulates the LAN configuration
// data.
type LanConfig map[string]string

// Response is a struct (intended for embedding) that represents response
// metadata returned by the EdgeOS API.
type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// FeatureResponse encapsulates the Response metadata and the Feature
type FeatureResponse struct {
	Response
	Feature Feature `json:"FEATURE"`
}

type Feature struct {
	Data PortForwards
	// Data       interface{}
	Definition interface{}

	Deletable string
	Success   string
}

func (c *Client) PortForwards() (*FeatureResponse, error) {
	res := &FeatureResponse{}
	err := c.FeatureFor(PortForwarding, res)
	return res, err
}
