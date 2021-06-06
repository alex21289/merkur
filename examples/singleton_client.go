package examples

import (
	"net/http"
	"time"

	"github.com/alex21289/merkur"
	"github.com/alex21289/merkur/mmime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() merkur.Client {
	headers := make(http.Header)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)

	client := merkur.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Alex-Computer").
		Build()
	return client
}
