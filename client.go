package merkur

import (
	"net/http"
	"sync"

	"github.com/alex21289/merkur/mcore"
)

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*mcore.Response, error)
	GetQuery(url string, params params, headers ...http.Header) (*mcore.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*mcore.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*mcore.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*mcore.Response, error)
	Delete(url string, headers ...http.Header) (*mcore.Response, error)
	Options(url string, headers ...http.Header) (*mcore.Response, error)
}

// Get
func (c *httpClient) Get(url string, headers ...http.Header) (*mcore.Response, error) {

	response, err := c.do(http.MethodGet, url, getHeaders(headers...), nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQuery does a get request with query parameters
func (c *httpClient) GetQuery(url string, p params, headers ...http.Header) (*mcore.Response, error) {

	if p != nil {
		qs := p.GetQueryString()
		url = url + qs
	}
	response, err := c.do(http.MethodGet, url, getHeaders(headers...), nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Post
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*mcore.Response, error) {
	response, err := c.do(http.MethodPost, url, getHeaders(headers...), body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Put
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*mcore.Response, error) {
	response, err := c.do(http.MethodPut, url, getHeaders(headers...), body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Patch
func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*mcore.Response, error) {
	response, err := c.do(http.MethodPatch, url, getHeaders(headers...), body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Delete
func (c *httpClient) Delete(url string, headers ...http.Header) (*mcore.Response, error) {
	response, err := c.do(http.MethodDelete, url, getHeaders(headers...), nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Options
func (c *httpClient) Options(url string, headers ...http.Header) (*mcore.Response, error) {
	response, err := c.do(http.MethodOptions, url, getHeaders(headers...), nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}
