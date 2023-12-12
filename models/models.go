package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"text;not null;default:null`
	Description string `json:"description" gorm:"text;not null;default:null`
	Status      string `json:"status" gorm:"text;not null;default:null`
}
