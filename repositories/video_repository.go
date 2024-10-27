package repositories

import (
	"database/sql"
	"errors"
	"ott-content-api/models"
)

type VideoRepository struct {
	db *sql.DB
}

// NewVideoRepository cria uma nova instância de VideoRepository
func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{db}
}

// Create cria um novo vídeo no banco de dados
func (repo *VideoRepository) Create(video *models.Video) error {
	_, err := repo.db.Exec("INSERT INTO videos (title, description, url) VALUES (?, ?, ?)", video.Title, video.Description, video.URL)
	return err
}

// Get busca um vídeo pelo ID
func (repo *VideoRepository) Get(id int) (*models.Video, error) {
	video := &models.Video{}
	row := repo.db.QueryRow("SELECT id, title, description, url FROM videos WHERE id = ?", id)

	err := row.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("video not found")
		}
		return nil, err
	}
	return video, nil
}

// Update atualiza um vídeo existente no banco de dados
func (repo *VideoRepository) Update(video *models.Video) error {
	_, err := repo.db.Exec("UPDATE videos SET title = ?, description = ?, url = ? WHERE id = ?", video.Title, video.Description, video.URL, video.ID)
	return err
}

// Delete remove um vídeo do banco de dados
func (repo *VideoRepository) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM videos WHERE id = ?", id)
	return err
}

// GetAll retorna todos os vídeos
func (repo *VideoRepository) GetAll() ([]*models.Video, error) {
	rows, err := repo.db.Query("SELECT id, title, description, url FROM videos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*models.Video
	for rows.Next() {
		video := &models.Video{}
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.URL); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}
