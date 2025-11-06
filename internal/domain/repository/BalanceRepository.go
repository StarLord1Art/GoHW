package repository

import "GoHW1/internal/domain/entity"

type BalanceRepository interface {
	CreateBalance(userId int64) (entity.Balance, error)
	GetBalance(userId int64) (int, error)
	TopUpBalance(userId int64, amountOfMoney int) (string, error)
	TransferToBalance(userIdFrom int64, userIdTo int64, amountOfMoney int) (string, error)
	InitializeBalanceStorage()
}
