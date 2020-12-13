package repositories

import (
	"github.com/jorgejr568/salary-go-api/db"
	"github.com/jorgejr568/salary-go-api/models"
	"github.com/jorgejr568/salary-go-api/requests"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryMongoRepository struct {
	collectionName string
	mongo          *db.MongoDbHelper
}

func (s SalaryMongoRepository) GetByUuid(uuid string) (*models.Salary, error) {
	var salary models.Salary
	err := s.mongo.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to mongodb")
		return nil, err
	}
	defer s.mongo.Disconnect()

	res := s.mongo.FindOneOnCollection(s.collectionName, struct {
		Uuid string `bson:"uuid"`
	}{Uuid: uuid})
	if res.Err() != nil {
		log.Error().Err(errors.WithStack(res.Err())).Msg("Salary not found!")
		return nil, res.Err()
	}

	resBody, _ := res.DecodeBytes()

	log.Info().Msg("Salary " + uuid + " found! " + resBody.String())
	err = res.Decode(&salary)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode salary from single result")
		return nil, err
	}

	return &salary, nil
}

func (s SalaryMongoRepository) Store(request *requests.SalaryStoreRequest) (*models.Salary, error) {
	err := s.mongo.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to mongodb")
		return nil, err
	}
	defer s.mongo.Disconnect()

	res, err := s.mongo.InsertOneOnCollection(s.collectionName, request)

	if err != nil {
		return nil, err
	}

	salary := &models.Salary{
		ID:               res.InsertedID.(primitive.ObjectID).Hex(),
		Uuid:             request.Uuid,
		Amount:           request.Amount,
		Description:      request.Description,
		CurrencyExchange: request.CurrencyExchange,
	}

	return salary, nil
}

func (s SalaryMongoRepository) Update(uuid string, request *requests.SalaryUpdateRequest) (*models.Salary, error) {
	return nil, nil
}

func NewSalaryMongoRepository() *SalaryMongoRepository {
	return &SalaryMongoRepository{
		collectionName: "salaries",
		mongo:          db.NewMongoDbHelper(),
	}
}
