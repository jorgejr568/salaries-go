package repositories

import (
	"github.com/jorgejr568/salary-go-api/models"
	"github.com/jorgejr568/salary-go-api/requests"
)

type SalaryRepositoryInterface interface {
	GetByUuid(uuid string) (*models.Salary, error)
	Store(request *requests.SalaryStoreRequest) (*models.Salary, error)
	Update(uuid string, request *requests.SalaryUpdateRequest) (*models.Salary, error)
}
