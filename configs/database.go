package configs

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task-management-api/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia%vJakarta", ENV.DB_HOST, ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_DATABASE, ENV.DB_PORT, "%2F")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database...")
	}

	db.AutoMigrate(&models.Task{})

	DB = db
	log.Info("Connected to database...")
}
