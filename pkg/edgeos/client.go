package edgeos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// A Client can interact with the EdgeOS REST API
type Client struct {
	Username, Password, Address string

	cli *http.Client
}

func (c *Client) endpoint(s string) string {
	return fmt.Sprintf("%s/api/edge/%s.json", c.Address, s)
}

// Login sets up an http session with the EdgeOS device using the supplied
// endpoint and credentials
func (c *Client) Login() error {
	v := url.Values{
		"username": []string{c.Username},
		"password": []string{c.Password},
	}

	res, err := c.cli.PostForm(c.Address+"/", v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Get returns some standard configuration information from EdgeOS
func (c *Client) Get() (map[string]interface{}, error) {
	res, err := c.cli.Get(c.endpoint("get"))

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var m map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&m)
	return m, err
}

// NewClient returns an initialized Client for interacting with an EdgeOS device
func NewClient(addr, username, password string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Username: username,
		Password: password,
		Address:  addr,
		cli: &http.Client{
			Transport: &csrfTransport{
				Referrer: addr,
			},
			Jar: jar,
		},
	}
	return c, nil
}
