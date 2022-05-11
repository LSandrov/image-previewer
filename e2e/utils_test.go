package e2e

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var c *http.Client
var timeout = 60 * time.Second

func init() {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c = &http.Client{
		Timeout:   timeout,
		Transport: customTransport,
	}
}

func getBaseUrl() string {
	return strings.TrimRight("http://127.0.0.1", "/")
}

func buildUrl(uri string) string {
	return fmt.Sprintf("%s/%s", getBaseUrl(), strings.TrimLeft(uri, "/"))
}

func readResponse(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}
