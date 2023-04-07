package repository

import (
	"github.com/gevs90/crypto-api/db/models"
	"gorm.io/gorm"
)

type TextRepository struct {
	db *gorm.DB
}

func NewTextRepository(db *gorm.DB) *TextRepository {
	return &TextRepository{db: db}
}

func (r *TextRepository) Create(text *models.Text) error {
	return r.db.Create(text).Error
}

func (r *TextRepository) Find(id int) (*models.Text, error) {
	var text models.Text
	err := r.db.Where("id = ?", id).First(&text).Error
	if err != nil {
		return nil, err
	}

	return &text, nil
}
