package migrations

import (
	"github.com/gevs90/crypto-api/db/models"
	"gorm.io/gorm"
)

func MigrateLogTable(db *gorm.DB) {
	db.AutoMigrate(&models.Log{})
}
