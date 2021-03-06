package models

import "time"

type Users struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	IsNew     uint      `json:"is_new"`
	Paket     Pakets    `json:"paket" gorm:"Foreignkey:id;association_foreignkey:id_paket;"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
