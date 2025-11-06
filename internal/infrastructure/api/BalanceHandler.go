package api

import (
	"GoHW1/internal/application/dto/request"
	"GoHW1/internal/application/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type BalanceHandler struct {
	balanceService usecase.BalanceService
}

func (bh *BalanceHandler) InitializeBalanceStorage() {
	bh.balanceService.InitializeBalanceStorage()
}

func (bh *BalanceHandler) CreateBalance(w http.ResponseWriter, r *http.Request) {
	var body request.BalanceRequestDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	res, err := bh.balanceService.CreateBalance(body.UserId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		return
	}
}

func (bh *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var body request.BalanceRequestDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	res, err := bh.balanceService.GetBalance(body.UserId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		return
	}
}

func (bh *BalanceHandler) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	var body request.BalanceReplenishmentRequestDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	res, err := bh.balanceService.TopUpBalance(body.UserId, body.AmountOfMoney)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		return
	}
}

func (bh *BalanceHandler) TransferToBalance(w http.ResponseWriter, r *http.Request) {
	var body request.BalanceTransferRequestDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	res, err := bh.balanceService.TransferToBalance(body.UserIdFrom, body.UserIdTo, body.AmountOfMoney)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		return
	}
}
