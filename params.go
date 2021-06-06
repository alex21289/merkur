package merkur

import "strings"

type params map[string][]string

func NewParams() params {
	p := params{}
	return p
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// The key is case insensitive;
func (p params) Add(key string, value string) {
	p[key] = append(p[key], value)
}

// Set sets the param entries associated with key to the
// single element value. It replaces any existing values
// associated with key. The key is case insensitive;
func (p params) Set(key string, value string) {
	p[key] = []string{value}
}

// Get gets the first value associated with the given key.
// It is case insensitive;
// If there are no values associated with the key, Get returns "".
func (p params) Get(key string) string {
	if p == nil {
		return ""
	}
	v := p[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

// Values returns all values associated with the given key.
// It is case insensitive;
// The returned slice is not a copy.
func (p params) Values(key string) []string {
	if p == nil {
		return nil
	}
	return p[key]
}

// Del deletes the values associated with key.
func (p params) Del(key string) {
	delete(p, key)
}

// GetQueryString returns a query string to append
// on a URL from a Params instance
func (p params) GetQueryString() string {
	var params []string
	for key, value := range p {
		for _, v := range value {
			params = append(params, key+"="+v)
		}
	}
	s := strings.Join(params, "&")
	return "?" + s
}
