package models

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"varchar(255)" json:"title"`
	Description string    `gorm:"varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
