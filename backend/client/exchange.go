package client

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/jorgejr568/salary-go-api/models"

	"github.com/jorgejr568/salary-go-api/cfg"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type ExchangeApiResponse struct {
	Rates map[string]float64 `json:"rates"`
	Date  string             `json:"date"`
	Base  string             `json:"base"`
}

type ExchangeApiClient struct {
	httpClient *fasthttp.Client
	baseUrl    string
}

func (e ExchangeApiClient) Exchange(salary *models.Salary) (*ExchangeClientResponse, error) {
	exchangeResponse, err := e.doGetForRates(salary)
	if err != nil {
		log.Error().Msg("Couldn't get rates for doing exchange")
		return nil, err
	}

	rate, exists := exchangeResponse.Rates[salary.CurrencyExchange.To]
	if !exists {
		log.Error().Msg("Couldn't find toCurrency on Rates map")
		return nil, fmt.Errorf("couldn't find toCurrency on Rates map")
	}

	toAmountFloat64 := float64(salary.Amount) * rate
	toAmountMinusTaxesFloat64 := toAmountFloat64 - (float64(salary.Amount) * salary.CurrencyExchange.RateTax)

	toAmount := int64(math.Round(toAmountFloat64))
	toAmountMinusTaxes := int64(math.Round(toAmountMinusTaxesFloat64)) - salary.CurrencyExchange.FixedTaxAmount

	return NewExchangeClientResponse(toAmount, toAmountMinusTaxes, rate), nil
}

func (e ExchangeApiClient) doGetForRates(salary *models.Salary) (*ExchangeApiResponse, error) {
	var exchangeResponse ExchangeApiResponse
	url := fmt.Sprintf("%s?base=%s&symbols=%s", e.baseUrl, salary.CurrencyExchange.From, salary.CurrencyExchange.To)
	r := e.doGetOnExchangeApi(url)

	if err := json.Unmarshal(r.Body(), &exchangeResponse); err != nil {
		log.Err(err).Msg("Couldn't unmarshal exchange response")
		return nil, err
	}

	return &exchangeResponse, nil

}
func (e ExchangeApiClient) doGetOnExchangeApi(uri string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)

	resp := fasthttp.AcquireResponse()
	e.httpClient.Do(req, resp)
	if resp.StatusCode() != fasthttp.StatusOK {
		log.Error().Int("status_code", resp.StatusCode()).Msg("Couldn't request ExchangeApiClient")
	}
	log.Info().RawJSON("responseBody", resp.Body()).Msg("Request ExchangeApi finished")
	return resp
}

func NewExchangeApiClient() *ExchangeApiClient {
	return &ExchangeApiClient{
		httpClient: &fasthttp.Client{},
		baseUrl:    cfg.CfgExchangeRatesApiBaseUrl(),
	}
}
