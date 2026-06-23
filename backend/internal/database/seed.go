package database

import (
	"backend/internal/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedData() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	// Jeżeli BD jest pusta to dodajemy testowe dane
	if count == 0 {
		log.Println("Baza jest pusta. Uruchamiam seeder...")

		// 1. Tworzenie użytkowników (SuperAdmin, Admin, Pilot)

		// Haszowanie haseł
		superAdminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		pilotPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

		// Dodawanie SuperAdmina
		superAdmin := models.User{
			Username: "superadmin",
			Password: string(superAdminPassword),
			Role:     "superadmin",
		}
		if err := DB.Create(&superAdmin).Error; err != nil {
			log.Println("Błąd podczas tworzenia konta SuperAdmin:", err)
		} else {
			log.Println("Pomyślnie utworzono SuperAdmin (Login: superadmin, Hasło: admin123)")
		}

		// Dodawanie zwykłego Admina
		DB.Create(&models.User{
			Username: "admin",
			Password: string(adminPassword),
			Role:     "admin",
		})
		log.Println("Pomyślnie utworzono Admin (Login: admin, Hasło: admin)")

		// Dodawanie Pilota
		DB.Create(&models.User{
			Username: "pilot",
			Password: string(pilotPassword),
			Role:     "pilot",
		})
		log.Println("Pomyślnie utworzono Pilot (Login: pilot, Hasło: 123)")

		// 2. Cennik (Tworzymy najpierw, aby uzyskać ich ID do relacji)
		price1 := models.Price{Type: "Cessna 152", PricePerMinute: 8.5}
		price2 := models.Price{Type: "Cessna 172", PricePerMinute: 11.0}
		price3 := models.Price{Type: "Piper PA-28", PricePerMinute: 10.0}

		DB.Create(&price1)
		DB.Create(&price2)
		DB.Create(&price3)

		// 3. Loty (Teraz tworzymy samoloty przypisując im odpowiednie PriceID)
		DB.Create(&models.Airplane{Model: "Cessna 152", Registration: "SP-KOR", Status: "available", PriceID: price1.ID})
		DB.Create(&models.Airplane{Model: "Cessna 172", Registration: "SP-LAL", Status: "maintenance", PriceID: price2.ID})
		DB.Create(&models.Airplane{Model: "Piper PA-28", Registration: "SP-FLY", Status: "available", PriceID: price3.ID})

		log.Println("Testowe dane zostały pomyślnie dodane do bazy!")
	}
}
