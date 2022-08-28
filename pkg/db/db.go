package db

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
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

// returns user name based on user id. user id is assumed to be in table.

func GetUserName(conn *gorm.DB, userId uint) string {
	user := User{}
	conn.First(&user, userId)
	return user.Name
}

// returns user bluetoothID given user id. user id is assumed to be in table.

func GetBluetoothID(conn *gorm.DB, userId uint) string {
	user := User{}
	conn.First(&user, userId)
	return user.BluetoothID
}

// returns user id on successful auth
func UserCheckCreds(conn *gorm.DB, email string, password string) (uint, error) {

	loginUser := User{}
	err := conn.
		Select("id", "password").
		Where("email = ?", email).
		First(&loginUser).
		Error
	if err != nil {
		return 0, err
	}

	// maybe move the bcrypt stuff into it's own function
	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(password))
	if err != nil {
		return 0, err
	}

	return loginUser.ID, nil
}

// check if username or email are already taken (true if not avaliable - possibly bad design)
func UserCheckRegistered(conn *gorm.DB, email string) (bool, error) {

	matches := []User{}
	err := conn.
		Where("email = ?", email).
		Find(&matches).
		Error
	if err != nil {
		return true, err
	}

	return len(matches) != 0, nil
}

/* a lot of these methods are very trivial - could just call db methods
 * directly in the client code
 */

func UserRegister(conn *gorm.DB, user *User) (uint, error) {

	err := conn.Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

// Get items related to userid

func GetList(conn *gorm.DB, userID uint) ([]Item, error) {
	items := []Item{}
	err := conn.
		Where("user_id = ?", userID).
		Find(&items).
		Error

	return items, err
}
