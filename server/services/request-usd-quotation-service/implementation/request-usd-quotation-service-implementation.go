package services

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/jeferreirajf/client-server-api-challenge/server/domain"
)

type RequestUsdQuotationServiceImpl struct{}

type EconomiaResponseDto struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func NewRequestUsdQuotationService() *RequestUsdQuotationServiceImpl {
	return &RequestUsdQuotationServiceImpl{}
}

func (r *RequestUsdQuotationServiceImpl) Request() (*domain.Quotation, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond) // 200 milliseconds
	defer cancel()

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	economiaResponse := EconomiaResponseDto{}

	err = json.Unmarshal(body, &economiaResponse)

	if err != nil {
		return nil, err
	}

	quotation := r.convertEconomiaResponseToQuotation(economiaResponse)

	select {
	case <-ctx.Done():
		return nil, errors.New("request service timeout")
	default:
		return quotation, nil
	}
}

func (r *RequestUsdQuotationServiceImpl) convertEconomiaResponseToQuotation(economiaResponse EconomiaResponseDto) *domain.Quotation {
	response := economiaResponse.USDBRL

	return domain.NewQuotation(
		response.Code,
		"1",
		response.Codein,
		response.Ask,
	)
}
