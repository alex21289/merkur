package merkur

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	maxIdelConnections int
	baseUrl            string
	client             *http.Client
	userAgent          string
	// baseUrl            string
}
type ClientBuilder interface {
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(*http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {

	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdelConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}
