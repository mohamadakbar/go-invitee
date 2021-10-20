package models

import "time"

func (Pakets) TableName() string {
	return "paket"
}

type Pakets struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug"`
	Name      string    `json:"nama_paket"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
