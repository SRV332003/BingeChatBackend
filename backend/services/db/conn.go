package db

import (
	"HangAroundBackend/config"
	"HangAroundBackend/models"
	"log"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(childLogger *zap.Logger) {
	// dbLogger := zapgorm2.New(childLogger)

	connectionString := config.GetEnv("DB_CONNECTION_STRING")
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		// Logger: dbLogger,
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	childLogger.Info("Connected to DB")
}

func Migrate(childLogger *zap.Logger) {
	err := Instance.AutoMigrate(&models.UserAuth{})
	if err != nil {
		childLogger.Error("Error migrating UserAuth", zap.Error(err))
		panic(err)	
	}
	err = Instance.AutoMigrate(&models.User{})
	if err != nil {
		childLogger.Error("Error migrating User", zap.Error(err))
		panic(err)
	}
	// write log to file

	childLogger.Info("Migrated tables")
}

func Disconnect() {
	db, err := Instance.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
