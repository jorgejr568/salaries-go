package main

import (
	"github.com/fasthttp/router"
	"github.com/jorgejr568/salary-go-api/cfg"
	"github.com/jorgejr568/salary-go-api/handlers"
	"github.com/lab259/cors"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func main() {
	cfg.SetupEnv()
	r := router.New()

	// Registering routes
	handlers.NewSalaryHandler().RegisterRoutes(r)

	port := cfg.CfgHttpPort()
	log.Info().Msg("Server running at http://localhost" + port)
	if err := fasthttp.ListenAndServe(port, cors.Default().Handler(r.Handler)); err != nil {
		log.Fatal().Err(err).Msg("could not ListenAndServe" + port)
	}
}
