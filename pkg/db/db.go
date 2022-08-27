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

func AddItem(conn *gorm.DB, item *Item) error {
	err := conn.Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

// returns user uid based on bluetooth id
func GetUserId(conn *gorm.DB, bluetoothId string) (uint, error) {

	user := User{}
	err := conn.
		Select("id").
		Where("bluetooth_id = ?", bluetoothId).
		First(&user).
		Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
