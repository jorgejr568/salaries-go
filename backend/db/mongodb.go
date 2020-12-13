package db

import (
	"context"
	"time"

	"github.com/jorgejr568/salary-go-api/cfg"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbHelper struct {
	client *mongo.Client
}

func (m MongoDbHelper) ctx() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	return ctx
}

func (m *MongoDbHelper) Connect() error {

	err := m.client.Ping(m.ctx(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not ping to mongodb")
	}

	return nil
}

func (m MongoDbHelper) Disconnect() {
	//m.client.Disconnect(m.ctx())
}

func (m MongoDbHelper) database() *mongo.Database {
	return m.client.Database(cfg.CfgMongoDatabase())
}
func (m MongoDbHelper) InsertOneOnCollection(collectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	return m.database().Collection(collectionName).InsertOne(m.ctx(), data)
}

func (m MongoDbHelper) FindOneOnCollection(collectionName string, filter interface{}) *mongo.SingleResult {
	return m.database().Collection(collectionName).FindOne(m.ctx(), filter)
}

func NewMongoDbHelper() *MongoDbHelper {
	client, err := GetClient()

	if err != nil {
		log.Fatal().Err(err).Msg("Could not create mongoDb client")
		return nil
	}
	return &MongoDbHelper{
		client: client,
	}
}

func GetClient() (*mongo.Client, error) {
	log.Info().Str("mongoUrl", cfg.CfgMongoUrl()).Msg("GetClient called")
	client, err := mongo.NewClient(
		options.Client().ApplyURI(cfg.CfgMongoUrl()))

	if err != nil {
		log.Fatal().Err(err).Msg("Could not create mongoDb client")
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	if err := client.Connect(ctx); err != nil {
		log.Error().Err(err).Msg("Could not connect to mongodb")
		return nil, err
	}

	return client, nil
}
