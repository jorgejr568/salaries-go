package requests

import (
	"time"

	"github.com/jorgejr568/salary-go-api/util"

	"github.com/google/uuid"
	"github.com/jorgejr568/salary-go-api/models"
)

type SalaryStoreRequest struct {
	Uuid             string                  `json:"uuid" bson:"uuid"`
	Amount           int64                   `json:"amount" bson:"amount"`
	Description      string                  `json:"description" bson:"description"`
	CurrencyExchange models.CurrencyExchange `json:"currency_exchange" bson:"currency_exchange"`
	CreatedAt        time.Time               `json:"created_at" bson:"created_at"`
}

type SalaryUpdateRequest struct {
	Amount           int64                   `json:"amount" bson:"amount"`
	Description      string                  `json:"description" bson:"description"`
	CurrencyExchange models.CurrencyExchange `json:"currency_exchange" bson:"currency_exchange"`
	UpdatedAt        time.Time               `json:"created_at" bson:"created_at"`
}

func NewSalaryStoreRequest(salary *models.Salary) *SalaryStoreRequest {
	return &SalaryStoreRequest{
		Uuid:             uuid.New().String(),
		Amount:           salary.Amount,
		Description:      salary.Description,
		CurrencyExchange: salary.CurrencyExchange,
		CreatedAt:        util.Now(),
	}
}

func NewSalaryUpdateRequest(salary *models.Salary) *SalaryUpdateRequest {
	return &SalaryUpdateRequest{
		Amount:           salary.Amount,
		Description:      salary.Description,
		CurrencyExchange: salary.CurrencyExchange,
		UpdatedAt:        util.Now(),
	}
}
