package app

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, response map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func isEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
