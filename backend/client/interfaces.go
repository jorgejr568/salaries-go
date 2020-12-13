package client

import "github.com/jorgejr568/salary-go-api/models"

type ExchangeClientInterface interface {
	Exchange(salary *models.Salary) (response *ExchangeClientResponse, err error)
}
