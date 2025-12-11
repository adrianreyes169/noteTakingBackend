package note

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ShowNotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		notes, err := ShowNotesService(db)

		if err != nil {
			http.Error(w, "Failed to retrieve notes"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(notes)
	}
}

func NotesByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "ID not provided", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(parts[2])

		if err != nil || id <= 0 {
			http.Error(w, "ID not valid", http.StatusBadRequest)
		}

		note, err := ShowNoteByID(db, id)

		if err != nil {
			http.Error(w, "Failed to retrieve note"+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	}
}

func CreateNotehandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var n Note
		err := json.NewDecoder(r.Body).Decode(&n)
		if err != nil {
			http.Error(w, "Invalid request Body", http.StatusBadRequest)
			return
		}

		id, err := CreateNote(db, n)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int64{"id": id})

	}
}

func DeleteNoteByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")

		if len(parts) != 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		_, err2 := DeleteNoteByIDService(db, id)

		if err2 != nil {
			http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)

	}
}
