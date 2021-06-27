package mcore

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
	Cookies    map[string]*http.Cookie
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
