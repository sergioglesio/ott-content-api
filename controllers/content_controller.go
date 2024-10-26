package controllers

import (
	"encoding/json"
	"net/http"
	"ott-content-api/models"
	"ott-content-api/repository"
	"ott-content-api/services"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type ContentController struct {
	Service *services.ContentService
}

func RegisterContentRoutes(r *mux.Router, db *gorm.DB) {
	contentRepo := repository.NewContentRepository(db)
	contentService := services.ContentService{Repository: contentRepo}
	controller := ContentController{Service: &contentService}

	r.HandleFunc("/contents", controller.CreateContent).Methods("POST")
	r.HandleFunc("/contents/{id}", controller.GetContentByID).Methods("GET")
	// Rota para Update e Delete aqui
}

func (c *ContentController) CreateContent(w http.ResponseWriter, r *http.Request) {
	var content models.Content
	json.NewDecoder(r.Body).Decode(&content)
	if err := c.Service.CreateContent(&content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(content)
}

func (c *ContentController) GetContentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	content, err := c.Service.GetContentByID(uint(id))
	if err != nil {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(content)
}
