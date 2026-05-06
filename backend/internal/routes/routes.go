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
		// Dostępne dla wszystkich (Pilotów i Adminów)
		protected.GET("/airplanes", controllers.GetAirplanes)
		protected.GET("/pricing", controllers.GetPrices)
		protected.GET("/reservations", controllers.GetReservations)
		protected.POST("/reservations", controllers.CreateReservation)

		// Strefa Administratora
		admin := protected.Group("/")
		admin.Use(middleware.RequireAdmin())
		{
			admin.PUT("/pricing", controllers.UpdatePrices)
		}
	}
}
