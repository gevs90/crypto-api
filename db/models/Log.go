package models

import "gorm.io/gorm"

type Log struct {
	Url    string `gorm:"not null"`
	Method string `gorm:"not null"`
	gorm.Model
}
