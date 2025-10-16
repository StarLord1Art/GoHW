package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var result struct {
	OutputString string `json:"outputString"`
}

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	res, err := client.Get("http://localhost:8081/version")
	if err != nil {
		log.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	fmt.Println(string(body))

	jsonData := map[string]string{
		"inputString": base64.StdEncoding.EncodeToString([]byte("yyy")),
	}
	data, err := json.Marshal(jsonData)
	if err != nil {
		log.Println(err)
	}
	res, err = client.Post("http://localhost:8081/decode", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result.OutputString)

	res, err = client.Get("http://localhost:8081/hard-op")
	if err != nil {
		fmt.Println(false)
		log.Println(err)
	}
	fmt.Print(true)
	fmt.Print(" ")
	fmt.Print(res.StatusCode)
}
