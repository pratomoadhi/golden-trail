package model

import "time"

type Transaction struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"index"`
	Amount      float64 `gorm:"not null"`
	Type        string  `gorm:"not null"` // income or expense
	Category    string
	Description string
	Date        time.Time
}
