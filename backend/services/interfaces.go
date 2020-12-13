package services

import (
	"github.com/jorgejr568/salary-go-api/models"
)

type SalaryServiceInterface interface {
	GetByUuid(uuid string) (*models.Salary, error)
	Store(salary *models.Salary) (*models.Salary, error)
	Update(uuid string, salary *models.Salary) (*models.Salary, error)
}

type ExchangeServiceInterface interface {
	Exchange(salary *models.Salary) (*models.SalaryExchanged, error)
}
