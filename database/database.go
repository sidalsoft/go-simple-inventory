package database

import (
	"errors"
	"fmt"

	"go-simple-inventory/models"
	"go-simple-inventory/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbName string) {

	var (
		databaseUser     string = utils.GetValue("DB_USER")
		databasePassword string = utils.GetValue("DB_PASSWORD")
		databaseHost     string = utils.GetValue("DB_HOST")
		databasePort     string = utils.GetValue("DB_PORT")
		databaseName     string = dbName
	)

	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	var err error

	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database")

	DB.AutoMigrate(&models.User{}, &models.Item{})
}

func SeedItem() (models.Item, error) {
	item, err := utils.CreateFaker[models.Item]()
	if err != nil {
		return models.Item{}, nil
	}

	DB.Create(&item)
	fmt.Println("Item seeded to the database")

	return item, nil
}

func SeedUser() (models.User, error) {
	user, err := utils.CreateFaker[models.User]()
	if err != nil {
		return models.User{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var inputUser models.User = models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: string(password),
	}

	DB.Create(&inputUser)
	fmt.Println("User seeded to the database")

	return user, nil
}

func CleanSeeders() {
	itemResult := DB.Exec("TRUNCATE items")
	userResult := DB.Exec("TRUNCATE users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	fmt.Println("Seeders are cleaned up successfully")
}
