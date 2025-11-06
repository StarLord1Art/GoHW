package usecase

import (
	"GoHW1/internal/application/dto/response"
	"GoHW1/internal/infrastructure/db"
	"errors"
)

type BalanceService struct {
	balanceRepository db.BalanceRepository
}

func (b *BalanceService) InitializeBalanceStorage() {
	b.balanceRepository.InitializeBalanceStorage()
}

func (b *BalanceService) CreateBalance(userId int64) (response.BalanceResponseDto, error) {
	balance, err := b.balanceRepository.CreateBalance(userId)
	if err != nil {
		return response.BalanceResponseDto{}, err
	}

	return response.BalanceResponseDto{UserId: balance.UserId, Balance: balance.Balance}, nil
}

func (b *BalanceService) GetBalance(userId int64) (response.BalanceResponseDto, error) {
	balance, err := b.balanceRepository.GetBalance(userId)
	if err != nil {
		return response.BalanceResponseDto{}, err
	}

	return response.BalanceResponseDto{UserId: userId, Balance: balance}, nil
}

func (b *BalanceService) TopUpBalance(userId int64, amountOfMoney int) (string, error) {
	if amountOfMoney <= 0 {
		return "", errors.New("amount of money must be positive")
	}

	res, err := b.balanceRepository.TopUpBalance(userId, amountOfMoney)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (b *BalanceService) TransferToBalance(userIdFrom int64, userIdTo int64, amountOfMoney int) (string, error) {
	if amountOfMoney <= 0 {
		return "", errors.New("amount of money must be positive")
	}
	if userIdFrom == userIdTo {
		return "", errors.New("you can not transfer to yourself")
	}

	res, err := b.balanceRepository.TransferToBalance(userIdFrom, userIdTo, amountOfMoney)
	if err != nil {
		return "", err
	}

	return res, nil
}
