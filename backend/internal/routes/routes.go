package routes

import (
	"backend/internal/controllers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Publiczne endpointy
	api.POST("/auth/login", controllers.Login)

	// Chronione endpointy (wymagany token)
	protected := api.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		// Dostępne dla wszystkich zalogowanych (Pilotów, Adminów, SuperAdminów)
		protected.GET("/airplanes", controllers.GetAirplanes)
		protected.GET("/pricing", controllers.GetPrices)
		protected.GET("/reservations", controllers.GetReservations)
		protected.POST("/reservations", controllers.CreateReservation)

		// Strefa Administratora (Zarządzanie Flotą ORAZ dodawanie użytkowników)
		admin := protected.Group("/")
		admin.Use(middleware.RequireAdmin()) // Puszcza admina i superadmina
		{
			admin.PUT("/pricing", controllers.UpdatePrices)
			admin.POST("/airplanes", controllers.CreateAirplane)
			admin.DELETE("/airplanes/:id", controllers.DeleteAirplane)
			// Admin może uderzać w ten endpoint, a wewnątrz kontrolera sprawdzamy rolę
			admin.POST("/users", controllers.CreateUser)
		}

		// (Grupa superadmin na razie może zostać pusta, jeśli nie masz innych specyficznych metod tylko dla niego)
	}
}
