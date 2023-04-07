package repository

import (
	"github.com/gevs90/crypto-api/db/models"
	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) Create(log *models.Log) error {
	return r.db.Create(log).Error
}
