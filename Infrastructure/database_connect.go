package Infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	// Auto migrate models
	err = DB.AutoMigrate(
	//&domain.Cart{},
	//&domain.Product{},
	//&domain.Order{},
	//&domain.Category{},
	//&domain.ProductCart{},
	//&domain.OrderDetail{},
	//&domain.ProductImage{},
	//&domain.ProductPromotion{},
	//&domain.Promotion{},
	//&domain.User{},
	//&domain.UserAddress{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	fmt.Println("Database connected and migrated!")
}

//func GetDB() *gorm.DB {
//	return DB
//}
