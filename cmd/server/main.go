package main

import (
	"log"
	"net/http"

	"first-api-go/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	app := app.NewApplication()

	r.Use(middleware.Logger)

	r.Get("/api/users/{id}", app.GetUserByID)
	r.Get("/api/users", app.GetAllUsers)
	r.Post("/api/users", app.CreateUser)
	r.Delete("/api/users/{id}", app.DeleteUser)
	r.Put("/api/users/{id}", app.UpdateUser)

	log.Println("Server running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
