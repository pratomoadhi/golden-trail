package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID uint      `json:"user_id"`                                                 // Foreign key
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"` // Optional eager-loading
	Amount float64   `json:"amount"`
	Type   string    `json:"type"` // "income" or "expense"
	Note   string    `json:"note"`
	Date   time.Time `json:"date"`
}

type TransactionInput struct {
	Amount float64 `json:"amount" binding:"required"`
	Type   string  `json:"type" binding:"required,oneof=income expense"`
	Note   string  `json:"note"`
	Date   string  `json:"date" binding:"required"` // ISO date (e.g., "2025-05-28")
}
