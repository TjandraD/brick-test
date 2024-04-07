package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(config Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
}
