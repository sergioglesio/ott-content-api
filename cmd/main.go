package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ott-content-api/models"
	"ott-content-api/repositories"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	dsn := "user-db-ott:mjhds@57fj92bn%@tcp(mysql:3306)/content"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	videoRepo := repositories.NewVideoRepository(db)

	r := mux.NewRouter()
	r.HandleFunc("/videos", createVideoHandler(videoRepo)).Methods("POST")
	r.HandleFunc("/videos/{id}", getVideoHandler(videoRepo)).Methods("GET")
	r.HandleFunc("/videos/{id}", updateVideoHandler(videoRepo)).Methods("PUT")
	r.HandleFunc("/videos/{id}", deleteVideoHandler(videoRepo)).Methods("DELETE")
	r.HandleFunc("/videos", getAllVideosHandler(videoRepo)).Methods("GET")

	fmt.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Manipuladores (handlers)
func createVideoHandler(repo *repositories.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var video models.Video
		if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := repo.Create(&video); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(video)
	}
}

func getVideoHandler(repo *repositories.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		video, err := repo.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(video)
	}
}

func updateVideoHandler(repo *repositories.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		var video models.Video
		if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		video.ID = id // Atualiza o ID para o que foi passado na URL
		if err := repo.Update(&video); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(video)
	}
}

func deleteVideoHandler(repo *repositories.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		if err := repo.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func getAllVideosHandler(repo *repositories.VideoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		videos, err := repo.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(videos)
	}
}
