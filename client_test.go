package sf

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"testing"
)

var c *Client

func TestNewClientWithOpt(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	c, err := NewClientWithOpts(WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func TestMain(m *testing.M) {
	var err error
	c, err = NewClientWithOpts()
	if err != nil {
		log.Fatal(err)
	}
	m.Run()
}
