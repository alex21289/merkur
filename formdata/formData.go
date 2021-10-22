package formdata

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type multiPartFormData struct {
	payload *bytes.Buffer
	writer  *multipart.Writer
	files   map[string]io.Writer
}

func (m *multiPartFormData) GetWriter() *multipart.Writer {
	return m.writer
}

func NewMultiPartForm() *multiPartFormData {
	payload := &bytes.Buffer{}
	f := make(map[string]io.Writer)
	return &multiPartFormData{
		payload: payload,
		writer:  multipart.NewWriter(payload),
		files:   f,
	}
}

func MarshalMulti(body interface{}) ([]byte, error) {
	d, ok := body.(*multiPartFormData)
	if ok {
		return d.GetPayload().Bytes(), nil
	}
	return nil, errors.New("cannot marshal formData body")
}

func (m *multiPartFormData) GetPayload() *bytes.Buffer {
	m.writer.Close()
	return m.payload
}

func (m *multiPartFormData) GetContentType() string {
	return m.writer.FormDataContentType()
}

func (m *multiPartFormData) AddField(name string, value string) error {
	return m.writer.WriteField(name, value)
}

func (m *multiPartFormData) AddFile(name string, filePath string) error {
	formFile, err := m.writer.CreateFormFile(name, filepath.Base(filePath))
	if err != nil {
		return err
	}

	m.files[name] = formFile

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(formFile, file)
	if err != nil {
		return err
	}

	return nil
}

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
