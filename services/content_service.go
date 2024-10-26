package services

import (
	"ott-content-api/models"
	"ott-content-api/repository"
)

type ContentService struct {
	Repository *repository.ContentRepository
}

func (s *ContentService) CreateContent(content *models.Content) error {
	return s.Repository.CreateContent(content)
}

func (s *ContentService) GetContentByID(id uint) (*models.Content, error) {
	return s.Repository.GetContentByID(id)
}

// Métodos Update e Delete aqui
