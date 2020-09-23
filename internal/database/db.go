package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/murilosrg/financial-api/config"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(config.Load().DB.Driver, config.Load().DB.Address)

	if err != nil {
		log.Fatalln("failed to connect database", err)
	}

	DB.Callback().Create().Remove("gorm:update_time_stamp")
}
