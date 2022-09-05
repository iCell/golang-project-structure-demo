package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewDBConn() *gorm.DB {
	cfg := &gorm.Config{Logger: logger.Default.LogMode(logger.Error), SkipDefaultTransaction: true}
	db, err := gorm.Open(postgres.Open("dsn"), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
