package main

import (
	"GoHW1/internal/infrastructure/api"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	userHandler := api.UserHandler{}
	userHandler.InitializeUserStorage()
	balanceHandler := api.BalanceHandler{}
	balanceHandler.InitializeBalanceStorage()

	mux.HandleFunc("POST /user/registration", userHandler.RegisterUser)

	mux.HandleFunc("POST /balance/creation", balanceHandler.CreateBalance)

	mux.HandleFunc("POST /balance/replenishment", balanceHandler.TopUpBalance)

	mux.HandleFunc("POST /balance/transfer", balanceHandler.TransferToBalance)

	mux.HandleFunc("GET /balance", balanceHandler.GetBalance)

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
