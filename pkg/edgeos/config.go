package edgeos

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
