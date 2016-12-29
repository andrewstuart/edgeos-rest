package edgeos

import (
	"net/http"
	"net/url"
)

type session struct {
	CStore map[string]string
	RTer   http.RoundTripper
}

func (c *session) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if c.CStore == nil {
		c.CStore = map[string]string{}
	}
	for _, ck := range cookies {
		c.CStore[ck.Name] = ck.Value
	}
}

//Cookies implements cookie store
func (c *session) Cookies(u *url.URL) []*http.Cookie {
	if c.CStore == nil || len(c.CStore) == 0 {
		c.CStore = map[string]string{}
	}

	cookies := []*http.Cookie{}
	for k, v := range c.CStore {
		cookies = append(cookies, &http.Cookie{Name: k, Value: v})
	}
	return cookies
}
