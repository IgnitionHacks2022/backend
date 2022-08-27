package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Migrate(conn *gorm.DB) error {
	return conn.Transaction(func(tsx *gorm.DB) error {

		err := tsx.AutoMigrate(
			&User{},
			&Item{},
		)
		if err != nil {
			return err
		}

		return nil
	})
}
