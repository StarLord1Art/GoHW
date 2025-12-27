package api

import (
	"GoHW1/internal/infrastructure/api/handler"
	"net/http"
)

func CreateRouting(userHandler *handler.UserHandler, balanceHandler *handler.BalanceHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/registration", userHandler.RegisterUser)

	mux.HandleFunc("POST /balance/creation", balanceHandler.CreateBalance)

	mux.HandleFunc("POST /balance/replenishment", balanceHandler.TopUpBalance)

	mux.HandleFunc("POST /balance/transfer", balanceHandler.TransferToBalance)

	mux.HandleFunc("GET /balance", balanceHandler.GetBalance)

	return mux
}
