package db

import (
	"GoHW1/internal/domain/entity"
	"errors"
)

type BalanceRepository struct {
	storage map[int64]entity.Balance
}

func (ubr *BalanceRepository) InitializeBalanceStorage() {
	ubr.storage = make(map[int64]entity.Balance)
}

func (ubr *BalanceRepository) CreateBalance(userId int64) (entity.Balance, error) {
	if _, ok := ubr.storage[userId]; ok {
		return entity.Balance{}, errors.New("balance already exists")
	}

	ubr.storage[userId] = entity.Balance{UserId: userId, Balance: 0}
	return ubr.storage[userId], nil
}

func (ubr *BalanceRepository) GetBalance(userId int64) (int, error) {
	return ubr.storage[userId].Balance, nil
}

func (ubr *BalanceRepository) TopUpBalance(userId int64, amountOfMoney int) (string, error) {
	entry, _ := ubr.storage[userId]
	entry.Balance += amountOfMoney
	ubr.storage[userId] = entry

	return "successful balance replenishment", nil
}

func (ubr *BalanceRepository) TransferToBalance(userIdFrom int64, userIdTo int64, amountOfMoney int) (string, error) {
	if ubr.storage[userIdFrom].Balance < amountOfMoney {
		return "", errors.New("not enough money on the balance")
	}

	entryFrom, _ := ubr.storage[userIdFrom]
	entryFrom.Balance -= amountOfMoney
	ubr.storage[userIdFrom] = entryFrom
	entryTo, _ := ubr.storage[userIdTo]
	entryTo.Balance += amountOfMoney
	ubr.storage[userIdTo] = entryTo

	return "successful transfer", nil
}
