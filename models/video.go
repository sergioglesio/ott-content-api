package models

type Video struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	ReleaseYear  int     `json:"release_year"`
	Director     string  `json:"director"`
	Stars        string  `json:"stars"`
	Genre        string  `json:"genre"`
	ThumbnailURL string  `json:"thumbnail_url"`
	URL          string  `json:"url"`
	Duration     int     `json:"duration"`
	Rating       float64 `json:"rating"`
}
