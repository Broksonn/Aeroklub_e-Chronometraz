package database

import (
	"fmt"
	"log"
	"os"

	"backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Ładujemy zmienne z pliku .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используем системные переменные")
	}

	// Budujemy ciąg połączenia ze zmiennych
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД: \n", err)
	}

	DB = db
	log.Println("Подключение к PostgreSQL успешно установлено")

	// Migracje i seeder
	DB.AutoMigrate(&models.User{}, &models.Airplane{}, &models.Price{}, &models.Reservation{})
	SeedData()
}
