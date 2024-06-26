package main

import (
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/controllers"
	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/go-chi/chi"
)

func main() {
	database.RunMigration()

	r := chi.NewRouter()

	r.Post("/poll/create", controllers.CreatePollController)
	r.Post("/vote/create", controllers.CreateVoteController)

	http.ListenAndServe(":4000", r)
}
