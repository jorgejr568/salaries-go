package services

import (
	"github.com/jorgejr568/salary-go-api/models"
	"github.com/jorgejr568/salary-go-api/repositories"
	"github.com/jorgejr568/salary-go-api/requests"
)

type SalaryServiceMongo struct {
	repository repositories.SalaryRepositoryInterface
}

func (s SalaryServiceMongo) GetByUuid(uuid string) (*models.Salary, error) {
	return s.repository.GetByUuid(uuid)
}

func (s SalaryServiceMongo) Store(salary *models.Salary) (*models.Salary, error) {
	request := requests.NewSalaryStoreRequest(salary)
	return s.repository.Store(request)
}

func (s SalaryServiceMongo) Update(uuid string, salary *models.Salary) (*models.Salary, error) {
	return nil, nil
}

func NewSalaryServiceMongo() *SalaryServiceMongo {
	return &SalaryServiceMongo{
		repository: repositories.NewSalaryMongoRepository(),
	}
}
