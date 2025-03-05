package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *Application) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, ok := app.Data[id]
	if !ok {
		RespondWithJSON(w, http.StatusNotFound, map[string]any{
			"message": "The user with the specified ID does not exist",
		})
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]any{
		"user": map[string]any{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"biography":  user.Biography,
			"id":         id,
		},
	})
}

func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []any{}
	for id, user := range app.Data {
		users = append(users, map[string]any{
			"id":         id,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"biography":  user.Biography,
		})
	}
	RespondWithJSON(w, http.StatusOK, map[string]any{
		"users": users,
	})
}

func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "unexpected internal server error",
			"error":   err.Error(),
		})
		return
	}

	if isEmptyString(body.FirstName) || isEmptyString(body.LastName) || isEmptyString(body.Biography) {
		RespondWithJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Please provide first_name, last_name, and biography",
		})
		return
	}

	_id := uuid.New().String()
	app.Data[_id] = body
	RespondWithJSON(w, http.StatusCreated, map[string]any{
		"user": map[string]any{
			"first_name": body.FirstName,
			"last_name":  body.LastName,
			"biography":  body.Biography,
			"id":         _id,
		},
	})
}

func (app *Application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, ok := app.Data[id]
	if !ok {
		RespondWithJSON(w, http.StatusNotFound, map[string]any{
			"message": "The user with the specified ID does not exist",
		})
		return
	}

	delete(app.Data, id)
	RespondWithJSON(w, http.StatusOK, map[string]any{
		"user": map[string]any{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"biography":  user.Biography,
			"id":         id,
		},
	})
}

func (app *Application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var body User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondWithJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "unexpected internal server error",
			"error":   err.Error(),
		})
		return
	}

	id := chi.URLParam(r, "id")
	user, ok := app.Data[id]
	if !ok {
		RespondWithJSON(w, http.StatusNotFound, map[string]any{
			"message": "The user with the specified ID does not exist",
		})
		return
	}

	if !isEmptyString(body.FirstName) {
		user.FirstName = body.FirstName
	}
	if !isEmptyString(body.LastName) {
		user.LastName = body.LastName
	}
	if !isEmptyString(body.Biography) {
		user.Biography = body.Biography
	}
	app.Data[id] = user

	RespondWithJSON(w, http.StatusOK, map[string]any{
		"user": map[string]any{
			"first_name": body.FirstName,
			"last_name":  body.LastName,
			"biography":  body.Biography,
			"id":         id,
		},
	})
}
