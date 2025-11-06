package request

type BalanceReplenishmentRequestDto struct {
	UserId        int64 `json:"userId"`
	AmountOfMoney int   `json:"amountOfMoney"`
}
