package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/jeferreirajf/client-server-api-challenge/server/domain"
)

type QuotationModel struct {
	ID           string `gorm:"primaryKey"`
	CurrencyFrom string
	ValueFrom    string
	CurrencyTo   string
	ValueTo      string
	CreatedAt    time.Time
}

func (QuotationModel) TableName() string {
	return "quotations"
}

func NewQuotationModel(q *domain.Quotation) *QuotationModel {
	return &QuotationModel{
		ID:           q.ID().String(),
		CurrencyFrom: q.CurrencyFrom(),
		ValueFrom:    q.ValueFrom(),
		CurrencyTo:   q.CurrencyTo(),
		ValueTo:      q.ValueTo(),
		CreatedAt:    q.CreatedAt(),
	}
}

func (q *QuotationModel) ToDomain() *domain.Quotation {
	id, _ := uuid.Parse(q.ID)

	return domain.NewQuotationWith(
		id,
		q.CurrencyFrom,
		q.ValueFrom,
		q.CurrencyTo,
		q.ValueTo,
		q.CreatedAt,
	)
}
