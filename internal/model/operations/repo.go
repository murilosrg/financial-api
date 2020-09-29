package operations

import (
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source=repo.go -destination=../mocks/mock_operation_repository.go -package=mocks

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
	if err := o.db.Create(&operationType).Error; err != nil {
		return err
	}

	return nil
}

func (o *OperationTypeRepository) Find(id int) (*OperationType, error) {
	operation := &OperationType{}
	if err := o.db.First(&operation, id).Error; err != nil {
		return nil, err
	}

	return operation, nil
}
