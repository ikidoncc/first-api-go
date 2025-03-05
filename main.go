package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func main() {
	r := chi.NewRouter()
	app := &application{data: make(map[string]user)}

	r.Use(middleware.Logger)

	// FindById - ok
	r.Get("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := app.data[id]
		if !ok {
			httpJsonResponse(w, http.StatusNotFound, map[string]any{
				"message": "The user with the specified ID does not exist",
			})
			return
		}
		httpJsonResponse(w, http.StatusOK, map[string]any{
			"user": map[string]any{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"biography":  user.Biography,
				"id":         id,
			},
		})
	})

	// FindAll - ok
	r.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		usersSlice := []any{}
		if len(app.data) > 0 {
			for id, user := range app.data {
				usersSlice = append(usersSlice, map[string]any{
					"first_name": user.FirstName,
					"last_name":  user.LastName,
					"biography":  user.Biography,
					"id":         string(id[:]),
				})
			}
		}

		httpJsonResponse(w, http.StatusOK, map[string]any{
			"users": usersSlice,
		})
	})

	// Create - ok
	r.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		var body user
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			httpJsonResponse(w, http.StatusInternalServerError, map[string]any{
				"message": "unexpected internal server error",
				"error":   err.Error(),
			})
			return
		}
		if isAnEmptyString(body.FirstName) || isAnEmptyString(body.LastName) || isAnEmptyString(body.Biography) {
			httpJsonResponse(w, http.StatusBadRequest, map[string]any{
				"message": "Please provide FirstName LastName and bio for the user",
			})
			return
		}
		_id := uuid.New().String()
		app.data[_id] = body
		user, ok := app.data[_id]
		if !ok {
			httpJsonResponse(w, http.StatusInternalServerError, map[string]any{
				"message": "There was an error while saving the user to the database",
			})
			return
		}
		httpJsonResponse(w, http.StatusCreated, map[string]any{
			"user": map[string]any{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"biography":  user.Biography,
				"id":         _id,
			},
		})
	})

	// Delete - ok
	r.Delete("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := app.data[id]
		if !ok {
			httpJsonResponse(w, http.StatusNotFound, map[string]any{
				"message": "The user with the specified ID does not exist",
			})
			return
		}
		delete(app.data, id)
		user, ok = app.data[id]
		if !ok {
			httpJsonResponse(w, http.StatusInternalServerError, map[string]any{
				"message": "The user could not be removed",
			})
			return
		}
		httpJsonResponse(w, http.StatusOK, map[string]any{
			"user": map[string]any{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"biography":  user.Biography,
				"id":         id,
			},
		})
	})

	// Update - ...
	r.Put("/api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var body user
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			httpJsonResponse(w, http.StatusInternalServerError, map[string]any{
				"message": "unexpected internal server error",
				"error":   err.Error(),
			})
			return
		}
		id := chi.URLParam(r, "id")
		user, ok := app.data[id]
		if !ok {
			httpJsonResponse(w, http.StatusNotFound, map[string]any{
				"message": "The user with the specified ID does not exist",
			})
			return
		}
		if !isAnEmptyString(body.FirstName) {
			user.FirstName = body.FirstName
		}
		if !isAnEmptyString(body.LastName) {
			user.LastName = body.LastName
		}
		if !isAnEmptyString(body.Biography) {
			user.Biography = body.Biography
		}
		app.data[id] = user
		httpJsonResponse(w, http.StatusOK, map[string]any{
			"user": map[string]any{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"biography":  user.Biography,
				"id":         id,
			},
		})
	})

	http.ListenAndServe(":3000", r)
}

/////////////////////////////////////
////// internal/model/user.go ///////
/////////////////////////////////////

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type application struct {
	data map[string]user
}

/////////////////////////////////////
////// Helpers/json.go //////////////
/////////////////////////////////////

func httpJsonResponse[K comparable, V any](w http.ResponseWriter, statusCode int, response map[K]V) {
	resp, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unexpected internal server error"))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func isAnEmptyString(s string) bool {
	s = strings.TrimSpace(s)
	return s == ""
}
