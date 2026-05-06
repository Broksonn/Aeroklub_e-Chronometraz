package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`    // Ukrywamy hasło podczas wysyłania JSON
	Role     string `json:"role"` // 'admin' lub 'pilot'
}
