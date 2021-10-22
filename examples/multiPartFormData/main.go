package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex21289/merkur"
	"github.com/alex21289/merkur/formdata"
)

func main() {

	client := merkur.NewBuilder().Build()

	data := formdata.NewMultiPartForm()
	data.AddField("display_name", "Foo-bar")
	data.AddField("document_type", "Invoice")
	data.AddField("layers", "Shopping,Amazon,Invoice,2021")
	data.AddField("document_date", time.Now().Format("2006-01-02"))
	data.AddFile("documentFile", "./test.txt")

	headers := make(http.Header)

	headers.Set("Authorization", "Bearer ")
	headers.Set("Content-Type", data.GetContentType())

	response, err := client.Post("http://localhost:8181/api/v1/documents/create", data, headers)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", response)
}
