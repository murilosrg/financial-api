package operations

import (
	"github.com/jinzhu/gorm"
	"github.com/murilosrg/financial-api/internal/database"
)

type IOperationTypeRepository interface {
	Create(operationType *OperationType) error
	Find(id int) (*OperationType, error)
}

type OperationTypeRepository struct {
	db *gorm.DB
}

func NewOperationTypeRepository(db *gorm.DB) IOperationTypeRepository {
	return &OperationTypeRepository{
		db: db,
	}
}

func (o *OperationTypeRepository) Create(operationType *OperationType) error {
	if err := database.DB.Create(&operationType).Error; err != nil {
		return err
	}

	return nil
}

func (o *OperationTypeRepository) Find(id int) (*OperationType, error) {
	operation := &OperationType{}
	if err := database.DB.First(&operation, id).Error; err != nil {
		return nil, err
	}

	return operation, nil
}
