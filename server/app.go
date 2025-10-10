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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("v1.0.0"))
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("POST /decode", func(w http.ResponseWriter, r *http.Request) {
		body, err1 := io.ReadAll(r.Body)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(r.Body)

		var data struct {
			InputString string `json:"inputString"`
		}
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}

		decodedString, err2 := base64.StdEncoding.DecodeString(data.InputString)
		if err2 != nil {
			log.Fatal(err2)
		}
		result := map[string]string{
			"outputString": string(decodedString),
		}
		w.Header().Set("Content-Type", "application/json")
		err3 := json.NewEncoder(w).Encode(result)
		if err3 != nil {
			log.Fatal(err3)
		}
	})

	mux.HandleFunc("GET /hard-op", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(10)+10) * time.Second)

		if rand.Intn(2)+1 == 1 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(rand.Intn(11) + 500)
		}
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	go func() {
		er := server.ListenAndServe()
		if er != nil {
			log.Println(er)
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
