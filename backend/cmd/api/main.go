package main

import (
	"log"

	"backend/internal/database"
	"backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Подключаемся к БД и запускаем миграции + сидер
	database.ConnectDB()

	// 2. Инициализируем роутер Gin
	r := gin.Default()

	// 3. Настройка CORS для связи с фронтом Vue (Vite обычно на порту 5173)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 4. Подключаем маршруты
	routes.SetupRoutes(r)

	// 5. Запуск сервера
	log.Println("Бэкенд запущен на http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
