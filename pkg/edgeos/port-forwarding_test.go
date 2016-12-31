package edgeos

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testJSON = `{"success":true,"FEATURE":{"data":{"wan":"eth0","hairpin-nat":"true","auto-firewall":"true","lans-config":[{"lan":"eth1"}],"rules-config":[{"description":"DNS","original-port":"53","protocol":"tcp_udp","forward-to-address":"192.168.16.5","forward-to-port":"53"},{"description":"OpenVPN Desktop","original-port":"1195","protocol":"udp","forward-to-address":"192.168.16.12","forward-to-port":"1195"},{"description":"Desk SSH","original-port":"2223","protocol":"tcp_udp","forward-to-address":"192.168.16.12","forward-to-port":"22"},{"description":"Vault","original-port":"8200","protocol":"tcp_udp","forward-to-address":"192.168.16.16","forward-to-port":"8200"},{"description":"OpenVPN HTPC","original-port":"1194","protocol":"udp","forward-to-address":"192.168.16.11","forward-to-port":"1194"},{"description":"SSH HTPC","original-port":"2222","protocol":"tcp_udp","forward-to-address":"192.168.16.11","forward-to-port":"22"},{"description":"Fire","original-port":"5431","protocol":"tcp_udp","forward-to-address":"192.168.16.12","forward-to-port":"5431"}]},"definition":{"lan":{"options":["eth2","eth0","eth1","imq0"],"other":"true"},"wan":{"options":["eth2","eth0","eth1","imq0"],"other":"true"}},"deletable":"1","success":"1"}}`

func TestDecode(t *testing.T) {
	asrt := assert.New(t)
	var res FeatureResponse

	asrt.NoError(json.NewDecoder(strings.NewReader(testJSON)).Decode(&res))

	asrt.True(res.Success)

	asrt.IsType(Feature{}, res.Feature)
	asrt.IsType(PortForwards{}, res.Feature.Data)
}
