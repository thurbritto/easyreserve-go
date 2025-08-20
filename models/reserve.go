package models

import (
	"time"
)

type Reserve struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"type:int" json:"user_id"`
	TableID     int       `gorm:"type:int" json:"table_id"`
	ReserveDate time.Time `json:"reserve_date"`
	Guests      int       `json:"guests"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
