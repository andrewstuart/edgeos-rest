# edgeos
--
    import "astuart.co/edgeos-rest/pkg/edgeos"

Package edgeos provides a Go client for interfacing with EdgeOS devices. It has
been developed and tested against the ERLite running v1.9.1 firmware. Thus far,
it serves primarily to expose high-level generic functionality as well as a
specific implementation for the port forwarding features.

## Usage

#### type Client

```go
type Client struct {
	Username, Password, Address string

	Path, Suffix, LoginEndpoint string
}
```

A Client can interact with the EdgeOS REST API

#### func  NewClient

```go
func NewClient(addr, username, password string) (*Client, error)
```
NewClient returns an initialized Client for interacting with an EdgeOS device

#### func (*Client) Feature

```go
func (c *Client) Feature(s Scenario) (Resp, error)
```
Feature takes an EdgeOS "Scenario" as an argument and returns a Resp
representing the JSON returned by the API.

#### func (*Client) FeatureFor

```go
func (c *Client) FeatureFor(s Scenario, out interface{}) error
```
FeatureFor takes a scenario and a pointer to a struct. The JSON response will be
deserialized into the `out` object, and any errors will be returned.

#### func (*Client) Get

```go
func (c *Client) Get() (Resp, error)
```
Get returns some standard configuration information from EdgeOS

#### func (*Client) GetJSON

```go
func (c *Client) GetJSON(endpoint string, data interface{}) (Resp, error)
```
GetJSON takes an endpoint and data for the POST body (or GET if `data` is nil)
and returns the Resp type that contains the data response from the endpoint.

#### func (*Client) JSONFor

```go
func (c *Client) JSONFor(endpoint string, data, out interface{}) error
```
JSONFor is a high-level method that takes an endpoint, a post body, and a
pointer to a struct into which the JSON should be decoded.

#### func (*Client) Login

```go
func (c *Client) Login() error
```
Login sets up an http session with the EdgeOS device using the supplied endpoint
and credentials

#### func (*Client) PortForwards

```go
func (c *Client) PortForwards() (*FeatureResponse, error)
```

#### func (*Client) SetFeature

```go
func (c *Client) SetFeature(s Scenario, data interface{}) (Resp, error)
```
SetFeature allows users to programmatically update features

#### func (*Client) SetFeatureFor

```go
func (c *Client) SetFeatureFor(s Scenario, data interface{}, out interface{}) error
```
SetFeatureFor takes a "Scenario", some data to send, and a pointer to an
interface into which the JSON response will be decoded.

#### type Config

```go
type Config struct {
	Success bool `json:"success"`
	GET     struct {
		Firewall struct {
			AllPing              string `json:"all-ping"`
			BroadcastPing        string `json:"broadcast-ping"`
			Ipv6ReceiveRedirects string `json:"ipv6-receive-redirects"`
			Ipv6SrcRoute         string `json:"ipv6-src-route"`
			IPSrcRoute           string `json:"ip-src-route"`
			LogMartians          string `json:"log-martians"`
			Name                 struct {
				LANLOCAL struct {
					DefaultAction string `json:"default-action"`
					Description   string `json:"description"`
					Rule          struct {
						Num1 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Destination struct {
								Group struct {
									AddressGroup string `json:"address-group"`
								} `json:"group"`
							} `json:"destination"`
							Log      string `json:"log"`
							Protocol string `json:"protocol"`
							Source   struct {
								Group interface{} `json:"group"`
							} `json:"source"`
							State struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"1"`
						Num2 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Destination struct {
								Group struct {
									AddressGroup string `json:"address-group"`
								} `json:"group"`
							} `json:"destination"`
							Log      string `json:"log"`
							Protocol string `json:"protocol"`
							Source   struct {
								Group interface{} `json:"group"`
							} `json:"source"`
							State struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"2"`
					} `json:"rule"`
				} `json:"LAN_LOCAL"`
				WANIN struct {
					DefaultAction    string      `json:"default-action"`
					Description      string      `json:"description"`
					EnableDefaultLog interface{} `json:"enable-default-log"`
					Rule             struct {
						Num20 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Log         string `json:"log"`
							Protocol    string `json:"protocol"`
							State       struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"20"`
						Num30 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Log         string `json:"log"`
							Protocol    string `json:"protocol"`
							Source      struct {
								Group interface{} `json:"group"`
							} `json:"source"`
							State struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"30"`
					} `json:"rule"`
				} `json:"WAN_IN"`
				WANLOCAL struct {
					DefaultAction    string      `json:"default-action"`
					Description      string      `json:"description"`
					EnableDefaultLog interface{} `json:"enable-default-log"`
					Rule             struct {
						Num10 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Log         string `json:"log"`
							Protocol    string `json:"protocol"`
							State       struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"10"`
						Num20 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Destination struct {
								Group struct {
									AddressGroup string `json:"address-group"`
								} `json:"group"`
							} `json:"destination"`
							Log      string `json:"log"`
							Protocol string `json:"protocol"`
						} `json:"20"`
						Num30 struct {
							Action      string `json:"action"`
							Description string `json:"description"`
							Log         string `json:"log"`
							Protocol    string `json:"protocol"`
							Source      struct {
								Group interface{} `json:"group"`
							} `json:"source"`
							State struct {
								Established string `json:"established"`
								Invalid     string `json:"invalid"`
								New         string `json:"new"`
								Related     string `json:"related"`
							} `json:"state"`
						} `json:"30"`
					} `json:"rule"`
				} `json:"WAN_LOCAL"`
			} `json:"name"`
			Options          interface{} `json:"options"`
			ReceiveRedirects string      `json:"receive-redirects"`
			SendRedirects    string      `json:"send-redirects"`
			SourceValidation string      `json:"source-validation"`
			SynCookies       string      `json:"syn-cookies"`
		} `json:"firewall"`
		Interfaces struct {
			Ethernet struct {
				Eth0 struct {
					Address     []string `json:"address"`
					Description string   `json:"description"`
					Duplex      string   `json:"duplex"`
					Firewall    struct {
						In struct {
							Name string `json:"name"`
						} `json:"in"`
						Local struct {
							Name string `json:"name"`
						} `json:"local"`
					} `json:"firewall"`
					Speed string `json:"speed"`
				} `json:"eth0"`
				Eth1 struct {
					Address     []string `json:"address"`
					Description string   `json:"description"`
					Duplex      string   `json:"duplex"`
					Firewall    struct {
						Local struct {
							Name string `json:"name"`
						} `json:"local"`
					} `json:"firewall"`
					Speed string `json:"speed"`
				} `json:"eth1"`
				Eth2 struct {
					Duplex string `json:"duplex"`
					Speed  string `json:"speed"`
				} `json:"eth2"`
			} `json:"ethernet"`
			Loopback struct {
				Lo interface{} `json:"lo"`
			} `json:"loopback"`
		} `json:"interfaces"`
		Service struct {
			DhcpServer struct {
				Disabled          string `json:"disabled"`
				HostfileUpdate    string `json:"hostfile-update"`
				SharedNetworkName struct {
					Internal struct {
						Authoritative string `json:"authoritative"`
						Subnet        struct {
							One9216816122 struct {
								DefaultRouter string   `json:"default-router"`
								DNSServer     []string `json:"dns-server"`
								DomainName    string   `json:"domain-name"`
								Lease         string   `json:"lease"`
								Start         struct {
									One9216816100 struct {
										Stop string `json:"stop"`
									} `json:"192.168.16.100"`
								} `json:"start"`
								StaticMapping struct {
									BRN30055CC581BE struct {
										IPAddress  string `json:"ip-address"`
										MacAddress string `json:"mac-address"`
									} `json:"BRN30055CC581BE"`
									R6300V2 struct {
										IPAddress  string `json:"ip-address"`
										MacAddress string `json:"mac-address"`
									} `json:"R6300v2"`
									Astuart struct {
										IPAddress  string `json:"ip-address"`
										MacAddress string `json:"mac-address"`
									} `json:"astuart"`
								} `json:"static-mapping"`
							} `json:"192.168.16.1/22"`
						} `json:"subnet"`
					} `json:"Internal"`
				} `json:"shared-network-name"`
				UseDnsmasq string `json:"use-dnsmasq"`
			} `json:"dhcp-server"`
			DNS struct {
				Forwarding struct {
					CacheSize string   `json:"cache-size"`
					ListenOn  []string `json:"listen-on"`
				} `json:"forwarding"`
			} `json:"dns"`
			Gui struct {
				HTTPPort     string `json:"http-port"`
				HTTPSPort    string `json:"https-port"`
				OlderCiphers string `json:"older-ciphers"`
			} `json:"gui"`
			Nat struct {
				Rule struct {
					Num5000 struct {
						Description       string `json:"description"`
						Log               string `json:"log"`
						OutboundInterface string `json:"outbound-interface"`
						Protocol          string `json:"protocol"`
						Source            struct {
							Group interface{} `json:"group"`
						} `json:"source"`
						Type string `json:"type"`
					} `json:"5000"`
				} `json:"rule"`
			} `json:"nat"`
			Snmp struct {
				Community struct {
					Public struct {
						Authorization string   `json:"authorization"`
						Network       []string `json:"network"`
					} `json:"public"`
				} `json:"community"`
				Contact  string `json:"contact"`
				Location string `json:"location"`
			} `json:"snmp"`
			SSH struct {
				Port            string `json:"port"`
				ProtocolVersion string `json:"protocol-version"`
			} `json:"ssh"`
		} `json:"service"`
		System struct {
			DomainName string `json:"domain-name"`
			HostName   string `json:"host-name"`
			Login      struct {
				User struct {
					Andrew struct {
						Authentication struct {
							EncryptedPassword string `json:"encrypted-password"`
							PlaintextPassword string `json:"plaintext-password"`
							PublicKeys        struct {
								AndrewAstuart struct {
									Key  string `json:"key"`
									Type string `json:"type"`
								} `json:"andrew@astuart"`
								AndrewDesktop struct {
									Key  string `json:"key"`
									Type string `json:"type"`
								} `json:"andrew@desktop"`
							} `json:"public-keys"`
						} `json:"authentication"`
						FullName string `json:"full-name"`
						Level    string `json:"level"`
					} `json:"andrew"`
				} `json:"user"`
			} `json:"login"`
			NameServer []string `json:"name-server"`
			Offload    struct {
				Hwnat string `json:"hwnat"`
				Ipsec string `json:"ipsec"`
				Ipv4  struct {
					Forwarding string `json:"forwarding"`
				} `json:"ipv4"`
				Ipv6 struct {
					Forwarding string `json:"forwarding"`
				} `json:"ipv6"`
			} `json:"offload"`
			Syslog struct {
				Global struct {
					Facility struct {
						All struct {
							Level string `json:"level"`
						} `json:"all"`
						Protocols struct {
							Level string `json:"level"`
						} `json:"protocols"`
					} `json:"facility"`
				} `json:"global"`
				Host struct {
					One9216816115140 struct {
						Facility struct {
							All struct {
								Level string `json:"level"`
							} `json:"all"`
						} `json:"facility"`
					} `json:"192.168.16.11:5140"`
					LogsAstuartCo5140 struct {
						Facility struct {
							All struct {
								Level string `json:"level"`
							} `json:"all"`
						} `json:"facility"`
					} `json:"logs.astuart.co:5140"`
				} `json:"host"`
			} `json:"syslog"`
			TimeZone        string `json:"time-zone"`
			TrafficAnalysis struct {
				Dpi    string `json:"dpi"`
				Export string `json:"export"`
			} `json:"traffic-analysis"`
		} `json:"system"`
		Vpn       interface{} `json:"vpn"`
		Protocols struct {
			Static struct {
				Route struct {
					One7231254124 struct {
						NextHop struct {
							One921681612 struct {
								Description string `json:"description"`
								Distance    string `json:"distance"`
							} `json:"192.168.16.12"`
						} `json:"next-hop"`
					} `json:"172.31.254.1/24"`
					One7231255124 struct {
						NextHop struct {
							One921681611 struct {
								Description string `json:"description"`
								Distance    string `json:"distance"`
							} `json:"192.168.16.11"`
						} `json:"next-hop"`
					} `json:"172.31.255.1/24"`
				} `json:"route"`
			} `json:"static"`
		} `json:"protocols"`
		TrafficControl interface{} `json:"traffic-control"`
	} `json:"GET"`
	SESSIONID string `json:"SESSION_ID"`
}
```


#### type Feature

```go
type Feature struct {
	Data PortForwards
	// Data       interface{}
	Definition interface{}

	Deletable string
	Success   string
}
```


#### type FeatureResponse

```go
type FeatureResponse struct {
	Response
	Feature Feature `json:"FEATURE"`
}
```

FeatureResponse encapsulates the Response metadata and the Feature

#### type LanConfig

```go
type LanConfig map[string]string
```

LanConfig is a simple string map that encapsulates the LAN configuration data.

#### type PortForward

```go
type PortForward struct {
	PortFrom    string `json:"original-port"`
	PortTo      string `json:"forward-to-port"`
	IPTo        string `json:"forward-to-address"`
	Protocol    string `json:"protocol"`
	Description string `json:"description"`
}
```

PortForward is a struct that represents a port forwarding rule

#### type PortForwards

```go
type PortForwards struct {
	AutoFirewall string        `json:"auto-firewall"`
	HairpinNAT   string        `json:"hairpin-nat"`
	WAN          string        `json:"wan"`
	Lans         []LanConfig   `json:"lans-config"`
	Rules        []PortForward `json:"rules-config"`
}
```

PortForwards is a struct that represents the response from the API for with the
list of port forwards, etc.

#### type Resp

```go
type Resp map[string]interface{}
```

Resp is the basic response type for the EdgeOS API. Higher-level methods will
tend to skip this type and return strongly-typed objects for specific endpoints.

#### type Response

```go
type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
```

Response is a struct (intended for embedding) that represents response metadata
returned by the EdgeOS API.

#### type Scenario

```go
type Scenario string
```

Scenario is just a string type to encourage the use of internal constants.

```go
const (
	PortForwarding Scenario = ".Port_Forwarding"
)
```
Common feature endpoints
