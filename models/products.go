package models

import "time"

type Products struct {
	ID		uint	`json:"id" gorm:"primary_key"`
	Code	string 	`json:"code"`
	Name	string	`json:"name"`
	Price	string	`json:"price"`
	CreatedAt	time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
