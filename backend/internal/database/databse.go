package database

import (
	"backend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// ВНИМАНИЕ: Замени password=root на свой пароль от Postgres!
	dsn := "host=localhost user=postgres password=root dbname=aeroclub port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к БД: \n", err)
	}

	DB = db
	log.Println("Подключение к PostgreSQL успешно установлено")

	// Автоматическое создание таблиц по моделям
	err = DB.AutoMigrate(
		&models.User{},
		&models.Airplane{},
		&models.Price{},
		&models.Reservation{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции: \n", err)
	}

	// Запускаем сидер
	SeedData()
}
