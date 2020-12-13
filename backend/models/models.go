package models

type CurrencyExchange struct {
	To             string  `json:"to" bson:"to"`
	From           string  `json:"from" bson:"from"`
	RateTax        float64 `json:"rate_tax" bson:"rate_tax"`
	FixedTaxAmount int64   `json:"fixed_tax_amount" bson:"fixed_tax_amount"`
}
type Salary struct {
	ID               string           `json:"id" bson:"_id"`
	Uuid             string           `json:"uuid" bson:"uuid"`
	Amount           int64            `json:"amount" bson:"amount"`
	Description      string           `json:"description,omitempty" bson:"description"`
	CurrencyExchange CurrencyExchange `json:"currency_exchange" bson:"currency_exchange"`
}

type SalaryExchanged struct {
	Salary  *Salary                 `json:"salary"`
	Rate    float64                 `json:"rate"`
	Amounts *SalaryExchangedAmounts `json:"amounts"`
}

type SalaryExchangedAmounts struct {
	Total      int64 `json:"total"`
	AfterTaxes int64 `json:"after_taxes"`
}

func NewSalaryExchanged(salary *Salary, rate float64, amounts *SalaryExchangedAmounts) *SalaryExchanged {
	return &SalaryExchanged{
		Salary:  salary,
		Rate:    rate,
		Amounts: amounts,
	}
}

func NewSalaryExchangedAmounts(total, afterTaxes int64) *SalaryExchangedAmounts {
	return &SalaryExchangedAmounts{
		Total:      total,
		AfterTaxes: afterTaxes,
	}
}
