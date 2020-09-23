package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GORMBase struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *GORMBase) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now())
	return nil
}
