package response

type BalanceResponseDto struct {
	UserId  int64 `json:"userId"`
	Balance int   `json:"balance"`
}
