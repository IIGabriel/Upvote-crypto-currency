package config

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	configDb := GetEnv("psql_settings")
	db, err := gorm.Open(postgres.Open(configDb), &gorm.Config{})

	if err != nil {
		zap.L().Panic("Could not connect to database", zap.Error(err))
	}

	return db
}

func CloseConnection(connection *gorm.DB) {
	db, err := connection.DB()
	if err != nil {
		zap.L().Panic("Could not close connection to database", zap.Error(err))
	}

	if err = db.Close(); err != nil {
		zap.L().Panic("Could not close connection to database", zap.Error(err))
	}
}
