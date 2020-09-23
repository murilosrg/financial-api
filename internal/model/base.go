package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type GORMBase struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time `json:"createAt"`
}

func (m *GORMBase) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}
