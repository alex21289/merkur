package mmock

import (
	"fmt"
	"net/http"

	"github.com/alex21289/merkur/mcore"
)

type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*mcore.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := mcore.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
