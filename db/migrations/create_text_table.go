package migrations

import (
	"github.com/gevs90/crypto-api/db/models"
	"gorm.io/gorm"
)

func MigrateTextTable(db *gorm.DB) {
	db.AutoMigrate(&models.Text{})
}
