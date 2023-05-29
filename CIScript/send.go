package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// set api url in the environment variable
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run send <file path> <api url>")
	}

	filePath := os.Args[1]
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	encodeContent := base64.StdEncoding.EncodeToString(fileContents)
	targetName := strings.Split(filepath.Base(filePath), ".")[0]

	formData := url.Values{}
	formData.Set("target_name", targetName)
	formData.Set("c_code", encodeContent)

	apiURL := os.Args[2]
	response, err := http.PostForm(apiURL, formData)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	apiResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(apiResponse))
}
