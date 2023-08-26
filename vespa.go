package vespa

import (
	"errors"
	"github.com/kinshotomoya/go-vespa/typedapi"
	"net/http"
	"time"
)

type TypedClient struct {
	BaseClient
	*typedapi.Api
}

func (client *TypedClient) Inform() {
}

func NewTypedClient(conf *Config) (*TypedClient, error) {
	if conf == nil {
		return nil, errors.New("conf is required")
	}
	httpClient, err := NewHttpClient(conf)
	if err != nil {
		return nil, err
	}

	client := TypedClient{
		BaseClient: BaseClient{
			Url:        conf.Url,
			HttpClient: httpClient,
		},
	}

	client.Api = typedapi.New(&client)

	return &client, nil
}

func NewHttpClient(conf *Config) (*http.Client, error) {
	if conf.Url == "" {
		return nil, errors.New("url is not determined")
	}

	var transport http.RoundTripper
	if conf.MaxActiveConnections == 0 || conf.MaxIdleConnections == 0 || conf.IdleConnTimeout == 0 {
		transport = http.DefaultTransport
	} else {
		transport = &http.Transport{
			MaxIdleConns:       conf.MaxIdleConnections,
			MaxConnsPerHost:    conf.MaxActiveConnections,
			DisableKeepAlives:  false,
			IdleConnTimeout:    time.Duration(conf.IdleConnTimeout) * time.Millisecond,
			DisableCompression: true,
		}
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(conf.TimeOut) * time.Millisecond,
	}, nil

}

type BaseClient struct {
	Url        string
	HttpClient *http.Client
}

type Config struct {
	HttpHeader           map[string]string
	TimeOut              int
	Url                  string
	MaxActiveConnections int
	MaxIdleConnections   int
	IdleConnTimeout      int
}
