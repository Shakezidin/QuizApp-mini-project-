package db

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	DOM "github.com/Shakezidin/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database(config config.Config) *gorm.DB {
	host := config.Host
	// user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=%s sslmode=%s", host, password, dbname, port, sslmode)

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to the database failed:", err)
	}

	// AutoMigrate all models
	err = DB.AutoMigrate(
		DOM.User{},
		
	)
	if err != nil {
		fmt.Println("error while migrating")
		return nil
	}

	return DB
}
