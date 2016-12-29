package edgeos

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

// type rtf func(*http.Request) (*http.Response, error)
type csrfTransport struct {
	Referrer string
	CSRF     string
	Debug    bool

	RoundTripper http.RoundTripper
}

func (r *csrfTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.RoundTripper == nil {
		r.RoundTripper = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}

	// req.Header.Set("Origin", "https://www.hackerrank.com")
	req.Header.Set("Referer", r.Referrer+"/")

	if r.CSRF != "" {
		req.Header.Set("X-CSRF-Token", r.CSRF)
	}

	res, err := r.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if res.Cookies() != nil {
		for _, cookie := range res.Cookies() {
			if cookie.Name == "X-CSRF-TOKEN" {
				r.CSRF = cookie.Value
			}
		}
	}

	if r.Debug {
		fmt.Println("Req")
		req.Write(os.Stdout)
		fmt.Println()

		fmt.Println("Res")
		res.Write(os.Stdout)
		fmt.Println()
	}

	return res, err
}
