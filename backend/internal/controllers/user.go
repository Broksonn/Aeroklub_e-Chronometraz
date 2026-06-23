package controllers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser obsługuje rejestrację nowego użytkownika przez admina / superadmina
func CreateUser(c *gin.Context) {
	// Pobieramy rolę zalogowanego użytkownika (twórcy), która została ustawiona w middleware RequireAuth
	creatorRole, exists := c.Get("userRole")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Brak autoryzacji"})
		return
	}

	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Role     string `json:"role" binding:"required,oneof=admin pilot superadmin"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe dane. Sprawdź format i upewnij się, że hasło ma min. 6 znaków"})
		return
	}

	// BLOKADA BEZPIECZEŃSTWA: Zwykły admin może tworzyć TYLKO konta pilotów
	if creatorRole == "admin" && input.Role != "pilot" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Brak uprawnień. Zwykły administrator może tworzyć tylko konta pilotów."})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd przetwarzania hasła"})
		return
	}

	newUser := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Użytkownik o takiej nazwie już istnieje"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pomyślnie utworzono nowego użytkownika"})
}
