package database

import (
	"backend/internal/models"
	"log"
)

func SeedData() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	// jezeli BD jest pusta to dodajemy testowe dane
	if count == 0 {
		log.Println("База пуста. Запускаем сидер...")

		// Users
		DB.Create(&models.User{Username: "admin", Password: "admin", Role: "admin"})
		DB.Create(&models.User{Username: "pilot", Password: "123", Role: "pilot"})

		// Loty
		DB.Create(&models.Airplane{Model: "Cessna 152", Registration: "SP-KOR", Status: "available"})
		DB.Create(&models.Airplane{Model: "Cessna 172", Registration: "SP-LAL", Status: "maintenance"})
		DB.Create(&models.Airplane{Model: "Piper PA-28", Registration: "SP-FLY", Status: "available"})

		// Cennik
		DB.Create(&models.Price{Type: "Cessna 152", PricePerMinute: 8.5})
		DB.Create(&models.Price{Type: "Cessna 172", PricePerMinute: 11.0})
		DB.Create(&models.Price{Type: "Piper PA-28", PricePerMinute: 10.0})

		log.Println("Testowe dane zostały pomyślnie dodane!")
	}
}
