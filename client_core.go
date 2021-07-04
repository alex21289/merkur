package merkur

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/alex21289/merkur/formdata"
	"github.com/alex21289/merkur/mcore"
	"github.com/alex21289/merkur/mmime"
	"github.com/alex21289/merkur/mmock"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 5 * time.Second
)

// do does a HTTP Request with the given parameters
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*mcore.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	request.Header = fullHeaders
	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Cookies
	cookies := make(map[string]*http.Cookie)
	for _, c := range response.Cookies() {
		cookies[c.Name] = c
	}

	finalResponse := mcore.Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
		Cookies:    cookies,
	}

	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() mcore.HttpClient {
	if mmock.MockupServer.IsMockServerEnabled() {
		return mmock.MockupServer.GetMockedClient()
	}
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		dialer := net.Dialer{
			Timeout: c.getConnectionTimeout(),
		}

		c.client = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext:           dialer.DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdelConnections > 0 {
		return c.builder.maxIdelConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case mmime.ContentTypeJson:
		return json.Marshal(body)

	case mmime.ContentTypeXml:
		return xml.Marshal(body)

	case mmime.ContentTypeXFormUrlencoded:
		return formdata.Marshal(body)

	case mmime.ContentTypeFormData:
		return formdata.Marshal(body)

	default:
		return json.Marshal(body)
	}
}
