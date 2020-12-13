package handlers

import (
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/jorgejr568/salary-go-api/models"
	"github.com/jorgejr568/salary-go-api/services"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type SalaryHandler struct {
	salaryService   services.SalaryServiceInterface
	exchangeService services.ExchangeServiceInterface
}

func (s SalaryHandler) RegisterRoutes(r *router.Router) {
	r.POST("/api/salary", s.Store)
	r.GET("/api/salary/{uuid}", s.Get)
	r.GET("/api/salary/{uuid}/exchange", s.GetExchange)
	r.PUT("/api/salary/{uuid}", s.Update)
}

func (s SalaryHandler) Store(ctx *fasthttp.RequestCtx) {
	var salaryRequest models.Salary
	if err := json.Unmarshal(ctx.PostBody(), &salaryRequest); err != nil {
		log.Error().Err(err).Msg("Could not unmarshal body from salary@store")
	}

	salary, err := s.salaryService.Store(&salaryRequest)

	if err != nil {
		log.Error().Err(err).Msg("Could not store salary")
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	respBody, _ := json.Marshal(salary)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusCreated)
	ctx.Write(respBody)
}

func (s SalaryHandler) Update(ctx *fasthttp.RequestCtx) {

}

func (s SalaryHandler) Get(ctx *fasthttp.RequestCtx) {
	uuid := s.uuidFromRoute(ctx)
	salary, err := s.salaryService.GetByUuid(uuid)

	if err != nil {
		log.Error().Err(err).Msg("Could not found salary")
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	respBody, _ := json.Marshal(salary)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(respBody)
}

func (s SalaryHandler) GetExchange(ctx *fasthttp.RequestCtx) {
	uuid := s.uuidFromRoute(ctx)

	salary, err := s.salaryService.GetByUuid(uuid)

	if err != nil {
		log.Error().Err(err).Msg("Could not found salary")
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	exchangeResponse, err := s.exchangeService.Exchange(salary)
	if err != nil {
		log.Error().Err(err).Msg("Could not exchange in api")
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	respBody, _ := json.Marshal(exchangeResponse)

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(respBody)
}

func (s SalaryHandler) uuidFromRoute(ctx *fasthttp.RequestCtx) string {
	return ctx.UserValue("uuid").(string)
}

func NewSalaryHandler() *SalaryHandler {
	return &SalaryHandler{
		salaryService:   services.NewSalaryServiceMongo(),
		exchangeService: services.NewExchangesRatesApiService(),
	}
}
