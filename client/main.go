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

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	res1, err1 := client.Get("http://localhost:8081/version")
	if err1 != nil {
		log.Fatal(err1)
	}
	body1, err2 := io.ReadAll(res1.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res1.Body)
	fmt.Println(string(body1))

	jsonData := map[string]string{
		"inputString": base64.StdEncoding.EncodeToString([]byte("yyy")),
	}
	data, e := json.Marshal(jsonData)
	if e != nil {
		log.Fatal(e)
	}
	res2, err3 := client.Post("http://localhost:8081/decode", "application/json", bytes.NewBuffer(data))
	if err3 != nil {
		log.Fatal(err3)
	}
	body2, err4 := io.ReadAll(res2.Body)
	if err4 != nil {
		log.Fatal(err4)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res2.Body)
	var result struct {
		OutputString string `json:"outputString"`
	}
	er := json.Unmarshal(body2, &result)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(result.OutputString)

	res3, err5 := client.Get("http://localhost:8081/hard-op")
	if err5 != nil {
		fmt.Println(false)
		log.Fatal(err5)
	}
	fmt.Print(true)
	fmt.Print(" ")
	fmt.Print(res3.StatusCode)
}
