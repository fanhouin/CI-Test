package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// set api url in the environment variable
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run send <file path> <api url>")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileField, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Fatal(err)
	}

	writer.Close()

	apiURL := os.Args[2]
	res, err := http.Post(apiURL, writer.FormDataContentType(), body)
	if err != nil {
		log.Fatal("Error while sending request")
	}
	defer res.Body.Close()

	msg, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(msg))
}
