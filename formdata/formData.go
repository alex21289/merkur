package formdata

import (
	"errors"
	"strings"
)

type formData map[string]string

// NewFormData creates a formData instance
func NewFormData() formData {
	fD := formData{}
	return fD
}

// Set sets the formData associated with key to the
// single element value. It replaces any existing values
// associated with key. The key is case insensitive;
func (fD formData) Set(key string, value string) {
	fD[key] = value
}

// Get gets the first value associated with the given key.
// It is case insensitive;
// If there are no values associated with the key, Get returns "".
func (fD formData) Get(key string) string {
	if fD == nil {
		return ""
	}
	return fD[key]
}

// Del deletes the values associated with key.
func (fD formData) Del(key string) {
	delete(fD, key)
}

// string returns a valid formdata string
func (fD formData) string() string {
	var formData []string
	for key, value := range fD {
		formData = append(formData, key+"="+value)
	}
	s := strings.Join(formData, "&")
	return s
}

// Marshal returns the formdata encoding of fd.
//
// Marshal can also handle fd as string and byte.
func Marshal(fd interface{}) ([]byte, error) {
	d, ok := fd.(formData)
	if ok {
		return []byte(d.string()), nil
	}
	ds, ok := fd.(string)
	if ok {
		return []byte(ds), nil
	}
	db, ok := fd.([]byte)
	if ok {
		return db, nil
	}

	return nil, errors.New("cannot marshal formData body")
}
