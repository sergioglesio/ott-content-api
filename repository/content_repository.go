package repository

import (
	"ott-content-api/models"

	"github.com/jinzhu/gorm"
)

type ContentRepository struct {
	DB *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{DB: db}
}

func (repo *ContentRepository) CreateContent(content *models.Content) error {
	return repo.DB.Create(content).Error
}

func (repo *ContentRepository) GetContentByID(id uint) (*models.Content, error) {
	var content models.Content
	err := repo.DB.First(&content, id).Error
	return &content, err
}

// Métodos Update e Delete aqui
