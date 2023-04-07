package models

import "gorm.io/gorm"

type Text struct {
	EncryptedText string `gorm:"text,not null"`
	EncryptionKey string `gorm:"text,not null"`
	gorm.Model
}
