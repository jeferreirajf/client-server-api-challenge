package domain

type QuotationGateway interface {
	Create(*Quotation) error
	FindById(id string) (*Quotation, error)
	List() ([]*Quotation, error)
}
