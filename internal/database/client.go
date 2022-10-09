package database

import (
	"fmt"

	"github.com/fabiotavarespr/golang-crud-rest-api/internal/entities"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot connect to DB - %s", err))
		panic("Cannot connect to DB")
	}
	logger.Info("Connected to Database...")
}
func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	logger.Info("Database Migration Completed...")
}
