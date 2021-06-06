package merkur

import (
	"net/http"

	"github.com/alex21289/merkur/mmime"
)

func getHeaders(headers ...http.Header) http.Header {
	requestHeaders := http.Header{}
	if len(headers) > 0 {
		for _, header := range headers {
			for key, val := range header {
				if len(val) > 0 {
					for _, v := range val {
						requestHeaders.Add(key, v)
					}
				}
			}
		}
		return requestHeaders
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers from the httpClient Instance
	for key, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	// Add custom headers to the result
	for key, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	// Add a User Agent to the headers if no userAgent is given in the result
	if c.builder.userAgent != "" {
		if result.Get(mmime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(mmime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
