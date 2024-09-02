package db

import (
	"HangAroundBackend/config"
	"HangAroundBackend/models"
	"log"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var Instance *gorm.DB
var dbError error

func Connect(childLogger *zap.Logger) {
	dbLogger := zapgorm2.New(childLogger)

	connectionString := config.GetEnv("DB_CONNECTION_STRING")
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: dbLogger,
	})
	if dbError != nil {
		childLogger.Panic("Error connecting to DB", zap.Error(dbError))
	}
	childLogger.Info("Connected to DB")
}

func Migrate(childLogger *zap.Logger) {

	err := Instance.AutoMigrate(&models.UserLogin{})
	if err != nil {
		childLogger.Panic("Error migrating User", zap.Error(err))
	}

	err = Instance.AutoMigrate(&models.UserInfo{})
	if err != nil {
		childLogger.Panic("Error migrating User", zap.Error(err))
	}

	err = Instance.AutoMigrate(&models.College{})
	if err != nil {
		childLogger.Panic("Error migrating College", zap.Error(err))
	}

	err = Instance.AutoMigrate(&models.Report{})
	if err != nil {
		childLogger.Panic("Error migrating Report", zap.Error(err))
	}

	childLogger.Info("Migrated tables")
}

func Disconnect() {
	db, err := Instance.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
