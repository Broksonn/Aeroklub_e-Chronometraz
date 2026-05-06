package models

import "time"

type Reservation struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	User       User      `json:"-" gorm:"foreignKey:UserID"`
	AirplaneID uint      `json:"airplane_id"`
	Airplane   Airplane  `json:"-" gorm:"foreignKey:AirplaneID"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
}
