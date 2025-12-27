package service

import "GoHW1/internal/application/dto/response"

type BalanceService interface {
	CreateBalance(userId int64) (response.BalanceResponseDto, error)
	GetBalance(userId int64) (response.BalanceResponseDto, error)
	TopUpBalance(userId int64, amountOfMoney int) (string, error)
	TransferToBalance(userIdFrom int64, userIdTo int64, amountOfMoney int) (string, error)
	InitializeBalanceStorage()
}
