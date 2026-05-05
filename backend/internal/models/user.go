package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`    // Скрываем пароль при отправке JSON
	Role     string `json:"role"` // 'admin' или 'pilot'
}
