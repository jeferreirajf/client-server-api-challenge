package domain

import (
	"time"

	"github.com/google/uuid"
)

type Quotation struct {
	id           uuid.UUID
	currencyFrom string
	valueFrom    string
	currencyTo   string
	valueTo      string
	createdAt    time.Time
}

func NewQuotation(currencyFrom string, valueFrom string, currencyTo string, valueTo string) *Quotation {
	quotation := Quotation{
		id:           uuid.New(),
		currencyFrom: currencyFrom,
		valueFrom:    valueFrom,
		currencyTo:   currencyTo,
		valueTo:      valueTo,
		createdAt:    time.Now(),
	}

	return &quotation
}

func NewQuotationWith(
	id uuid.UUID,
	currencyFrom string,
	valueFrom string,
	currencyTo string,
	valueTo string,
	createdAt time.Time,
) *Quotation {
	quotation := Quotation{
		id:           id,
		currencyFrom: currencyFrom,
		valueFrom:    valueFrom,
		currencyTo:   currencyTo,
		valueTo:      valueTo,
		createdAt:    createdAt,
	}

	return &quotation
}

func (q *Quotation) ID() uuid.UUID {
	return q.id
}

func (q *Quotation) CurrencyFrom() string {
	return q.currencyFrom
}

func (q *Quotation) ValueFrom() string {
	return q.valueFrom
}

func (q *Quotation) CurrencyTo() string {
	return q.currencyTo
}

func (q *Quotation) ValueTo() string {
	return q.valueTo
}

func (q *Quotation) CreatedAt() time.Time {
	return q.createdAt
}

func (q *Quotation) ToString() string {
	return "Quotation: " + q.id.String() + "\nFrom: " + q.currencyFrom + " " + q.valueFrom + "\nTo: " + q.currencyTo + " " + q.valueTo
}
