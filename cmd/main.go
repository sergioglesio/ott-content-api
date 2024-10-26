package main

import (
	"log"
	"net/http"
	"ott-content-api/config"
	"ott-content-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Carregar configurações e conectar ao banco de dados
	db := config.InitDB()
	defer db.Close()

	// Inicializar o roteador e configurar rotas
	r := mux.NewRouter()
	controllers.RegisterContentRoutes(r, db)

	// Iniciar servidor
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
