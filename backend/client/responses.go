package client

type ExchangeClientResponse struct {
	Amount           int64
	AmountMinusTaxes int64
	Rate             float64
}

func NewExchangeClientResponse(amount, amountMinusTaxes int64, rateUsed float64) *ExchangeClientResponse {
	return &ExchangeClientResponse{
		Amount:           amount,
		AmountMinusTaxes: amountMinusTaxes,
		Rate:             rateUsed,
	}
}
