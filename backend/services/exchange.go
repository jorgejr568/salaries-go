package services

import (
	"github.com/jorgejr568/salary-go-api/client"
	"github.com/jorgejr568/salary-go-api/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type ExchangesRatesApiService struct {
	client client.ExchangeClientInterface
}

func (e ExchangesRatesApiService) Exchange(salary *models.Salary) (*models.SalaryExchanged, error) {
	response, err := e.client.Exchange(salary)
	if err != nil {
		log.Error().Err(errors.WithStack(err)).Msg("Exchange service error on Exchange method")
		return nil, err
	}

	return models.NewSalaryExchanged(
		salary, response.Rate, models.NewSalaryExchangedAmounts(response.Amount, response.AmountMinusTaxes)), nil
}

func NewExchangesRatesApiService() *ExchangesRatesApiService {
	return &ExchangesRatesApiService{
		client: client.NewExchangeApiClient(),
	}
}
