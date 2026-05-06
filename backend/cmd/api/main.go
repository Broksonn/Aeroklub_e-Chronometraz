package main

import (
	"log"

	"backend/internal/database"
	"backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Łączymy się z bazą danych i uruchamiamy migracje + seeder
	database.ConnectDB()

	// 2. Inicjalizujemy router Gin
	r := gin.Default()

	// 3. Konfiguracja CORS do komunikacji z frontendem Vue (Vite zazwyczaj na porcie 5173)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 4. Podłączamy trasy
	routes.SetupRoutes(r)

	// 5. Uruchomienie serwera
	log.Println("Backend uruchomiony na http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Blad podczas uruchamiania serwera: ", err)
	}
}
