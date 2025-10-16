package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func version(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("v1.0.0"))
	if err != nil {
		log.Println(err)
	}
}

var data struct {
	InputString string `json:"inputString"`
}

func decode(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	decodedString, err := base64.StdEncoding.DecodeString(data.InputString)
	if err != nil {
		log.Println(err)
	}
	result := map[string]string{
		"outputString": string(decodedString),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
	}
}

func hardOperation(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(10)+10) * time.Second)

	if rand.Intn(2)+1 == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(rand.Intn(11) + 500)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /version", version)

	mux.HandleFunc("POST /decode", decode)

	mux.HandleFunc("GET /hard-op", hardOperation)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	<-osSignals
	log.Println("Graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
