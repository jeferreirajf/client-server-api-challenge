package services

import (
	"github.com/jeferreirajf/client-server-api-challenge/server/domain"
)

type RequestUsdQuotationService interface {
	Request() (*domain.Quotation, error)
}
