package main

import (
	"log"
	"net/http"

	"github.com/alex21289/merkur"
	"github.com/alex21289/merkur/formdata"
	"github.com/alex21289/merkur/mmime"
)

func main() {

	client := merkur.NewBuilder().Build()

	formData := formdata.NewFormData()
	formData.Set("Moie", "1234")
	formData.Set("Hallo", "sadasds")

	// formData := "test=1234&token=12345231nmdasr"
	headers := make(http.Header)

	headers.Set("Content-Type", mmime.ContentTypeXFormUrlencoded)
	response, err := client.Post("http://localhost:9191/api", formData, headers)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", response)
}
