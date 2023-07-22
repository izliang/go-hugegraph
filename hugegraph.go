package hugegraph

import (
	"errors"
	"fmt"
	v1 "github.com/izliang/go-hugegraph/hgapi/v1"
	v3 "github.com/izliang/go-hugegraph/hgapi/v3"
	"github.com/izliang/go-hugegraph/hgtransport"
	"net"
	"net/http"
	"net/url"
	"os"
)

// ClientTypeEnum 客户端类型枚举
type ClientTypeEnum string

// DefaultClientType 客户端类型-基础
const CommonClientType ClientTypeEnum = ""

// StarDefaultClientType 客户端类型-star版
const StarDefaultClientType ClientTypeEnum = "star"

// Config 配置类型
type Config struct {
	Host       string
	Port       int
	GraphSpace string
	Graph      string
	Username   string
	Password   string

	ClientType ClientTypeEnum     // Client Type
	Transport  http.RoundTripper  // The HTTP transport object.
	Logger     hgtransport.Logger // The logger object.
}

type CommonClient struct {
	*v1.APIV1
	Transport hgtransport.Interface
	Graph     string
}

type StarClient struct {
	*v3.APIV3
	Transport hgtransport.Interface
	Graph     string
}

func NewDefaultCommonClient() (*CommonClient, error) {
	return NewCommonClient(Config{
		Host:     "127.0.0.1",
		Port:     8080,
		Graph:    "1",
		Username: "1",
		Password: "1",
		Logger: &hgtransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})
}

func NewDefaultStarClient() (*StarClient, error) {
	return NewStarClient(Config{
		Host:     "127.0.0.1",
		Port:     8080,
		Graph:    "1",
		Username: "1",
		Password: "1",
		Logger: &hgtransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	})
}

func NewCommonClient(cfg Config) (*CommonClient, error) {

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
	if cfg.ClientType != CommonClientType {
		return nil, errors.New("cannot create client: NewCommonClient only supported commonClient")
	}

	tp := hgtransport.New(hgtransport.Config{
		URL: &url.URL{
			Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Scheme: "http",
		},
		Username:  cfg.Username,
		Password:  cfg.Password,
		Graph:     cfg.Graph,
		Transport: cfg.Transport,
		Logger:    cfg.Logger,
	})

	return &CommonClient{
		APIV1:     v1.New(tp),
		Transport: tp,
		Graph:     cfg.Graph,
	}, nil
}

func NewStarClient(cfg Config) (*StarClient, error) {

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
	if cfg.ClientType != StarDefaultClientType {
		return nil, errors.New("cannot create client: NewStarClient only supported starClient")
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
		Transport:  cfg.Transport,
		Logger:     cfg.Logger,
	})

	return &StarClient{Transport: tp, APIV3: v3.New(tp), Graph: cfg.Graph}, nil
}
