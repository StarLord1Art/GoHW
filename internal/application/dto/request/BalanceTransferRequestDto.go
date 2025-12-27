package request

type BalanceTransferRequestDto struct {
	UserIdFrom    int64 `json:"userIdFrom"`
	UserIdTo      int64 `json:"userIdTo"`
	AmountOfMoney int   `json:"amountOfMoney"`
}
