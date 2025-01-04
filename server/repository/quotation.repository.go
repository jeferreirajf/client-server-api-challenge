package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jeferreirajf/client-server-api-challenge/server/domain"
	repository "github.com/jeferreirajf/client-server-api-challenge/server/repository/models"
	"gorm.io/gorm"
)

type QuotationRepository struct {
	db *gorm.DB
}

func NewQuotationRepository(db gorm.DB) *QuotationRepository {

	db.AutoMigrate(&repository.QuotationModel{})

	return &QuotationRepository{
		db: &db,
	}
}

func (qr *QuotationRepository) Create(quotation *domain.Quotation) error {
	aModel := repository.NewQuotationModel(quotation)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond) // 10 milliseconds
	defer cancel()

	result := qr.db.WithContext(ctx).Create(aModel)

	if result.Error != nil {
		return result.Error
	}

	select {
	case <-ctx.Done():
		return errors.New("repository timeout")
	default:
		return nil
	}
}

func (qr *QuotationRepository) FindById(id string) (*domain.Quotation, error) {
	var aModel repository.QuotationModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second/1000) // 10 milliseconds
	defer cancel()

	result := qr.db.WithContext(ctx).First(&aModel, id)

	if result.Error != nil {
		return nil, result.Error
	}

	aQuotation := aModel.ToDomain()

	return aQuotation, nil
}

func (qr *QuotationRepository) List() ([]*domain.Quotation, error) {
	var aModels []*repository.QuotationModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second/1000) // 10 milliseconds
	defer cancel()

	result := qr.db.WithContext(ctx).Find(&aModels)

	if result.Error != nil {
		return nil, result.Error
	}

	var quotations []*domain.Quotation

	for _, aModel := range aModels {
		quotation := aModel.ToDomain()
		quotations = append(quotations, quotation)
	}

	return quotations, nil
}
