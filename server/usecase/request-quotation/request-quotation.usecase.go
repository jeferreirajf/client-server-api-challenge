package usecase

import (
	"github.com/jeferreirajf/client-server-api-challenge/server/domain"
	services "github.com/jeferreirajf/client-server-api-challenge/server/services/request-usd-quotation-service"
)

type RequestQuotationInputDto struct{}

type RequestQuotationOutputDto struct {
	Quotation struct {
		ID           string `json:"id"`
		CurrencyFrom string `json:"currency_from"`
		ValueFrom    string `json:"value_from"`
		CurrencyTo   string `json:"currency_to"`
		ValueTo      string `json:"value_to"`
		CreatedAt    string `json:"created_at"`
	} `json:"quotation"`
}

type RequestQuotationUsecase struct {
	quotationGateway        domain.QuotationGateway
	requestQuotationService services.RequestUsdQuotationService
}

func NewRequestQuotationUsecase(quotationGateway domain.QuotationGateway, requestQuotationService services.RequestUsdQuotationService) *RequestQuotationUsecase {
	return &RequestQuotationUsecase{
		quotationGateway:        quotationGateway,
		requestQuotationService: requestQuotationService,
	}
}

func (u *RequestQuotationUsecase) Execute(inputDto *RequestQuotationInputDto) (*RequestQuotationOutputDto, error) {
	quotation, err := u.requestQuotationService.Request()

	if err != nil {
		return nil, err
	}

	err = u.quotationGateway.Create(quotation)

	if err != nil {
		return nil, err
	}

	outputDto := &RequestQuotationOutputDto{}
	outputDto.Quotation.ID = quotation.ID().String()
	outputDto.Quotation.CurrencyFrom = quotation.CurrencyFrom()
	outputDto.Quotation.ValueFrom = quotation.ValueFrom()
	outputDto.Quotation.CurrencyTo = quotation.CurrencyTo()
	outputDto.Quotation.ValueTo = quotation.ValueTo()
	outputDto.Quotation.CreatedAt = quotation.CreatedAt().String()

	return outputDto, nil
}
