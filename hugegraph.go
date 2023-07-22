package hugegraph

import (
	"errors"
	"fmt"
	"hugegraph/hgapi"
	"hugegraph/hgtransport"
	"net"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	Host       string
	Port       int
	GraphSpace string
	Graph      string
	Username   string
	Password   string

	Transport http.RoundTripper  // The HTTP transport object.
	Logger    hgtransport.Logger // The logger object.
}

type Client struct {
	*hgapi.API
	Transport hgtransport.Interface
	Graph     string
}

func NewDefaultClient() (*Client, error) {
	return NewClient(Config{
		Host:       "10.41.58.84",
		Port:       8084,
		GraphSpace: "baikegs",
		Graph:      "lemma_test",
		Username:   "baike_dp",
		Password:   "8221a0515d30c988",
		Logger: &hgtransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})
}

func NewClient(cfg Config) (*Client, error) {

	if len(cfg.Host) < 3 {
		return nil, errors.New("cannot create client: host length error")
	}
	address := net.ParseIP(cfg.Host)
	if address == nil {
		return nil, errors.New("cannot create client: host is format error")
	}
	if cfg.Port < 1 || cfg.Port > 65535 {
		return nil, errors.New("cannot create client: port is error")
	}

	tp := hgtransport.New(hgtransport.Config{
		URL: &url.URL{
			Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Scheme: "http",
		},
		Username:   cfg.Username,
		Password:   cfg.Password,
		GraphSpace: cfg.GraphSpace,
		Graph:      cfg.Graph,

		Transport: cfg.Transport,
		Logger:    cfg.Logger,
	})

	return &Client{Transport: tp, API: hgapi.New(tp)}, nil
}
