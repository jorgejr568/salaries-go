package cfg

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	envName                 = "ENV_NAME"
	mongoUrl                = "MONGO_URL"
	mongoDatabase           = "MONGO_DATABASE"
	exchangeRatesApiBaseUrl = "EXCHANGE_RATES_API_BASE_URL"
	defaultTimezone         = "DEFAULT_TIMEZONE"
	httpPort                = "HTTP_PORT"
)

func SetupEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Panic().Err(errors.WithStack(err)).Msg("Could not load environment variables")
		return err
	}

	return nil
}

func withFallback(env, fallback string) string {
	if v, e := os.LookupEnv(env); e {
		return v
	}
	return fallback
}

func CfgEnvName() string {
	return withFallback(envName, "development")
}

func CfgMongoUrl() string {
	return withFallback(mongoUrl, "mongodb://root:secret@db:27017")
}

func CfgMongoDatabase() string {
	return withFallback(mongoDatabase, "salaryAPI")
}

func CfgExchangeRatesApiBaseUrl() string {
	return withFallback(exchangeRatesApiBaseUrl, "https://api.exchangeratesapi.io/latest")
}

func CfgDefaultTimezone() string {
	return withFallback(defaultTimezone, "America/Sao_Paulo")
}

func CfgHttpPort() string {
	return withFallback(httpPort, ":8080")
}
