package operations

type OperationType struct {
	ID          int    `gorm:"primaryKey"`
	Description string `json:"description" gorm:"not null;type:varchar(50)"`
	IsDebit     bool   `json:"is_debit" gorm:"not null;column:debit"`
}

func (OperationType) TableName() string {
	return "operation_type"
}
