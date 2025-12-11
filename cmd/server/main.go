package main

import (
	"log"
	"net/http"

	"NoteTakingBackend/internal/db"
	"NoteTakingBackend/note"
)

func main() {
	database := db.DBConnection()
	defer database.Close()

	http.HandleFunc("/notes", note.ShowNotesHandler(database))
	http.HandleFunc("/notesbyid/", note.NotesByIDHandler(database))
	http.HandleFunc("/create", note.CreateNotehandler(database))
	http.HandleFunc("/delete/", note.DeleteNoteByIDHandler(database))

	log.Println("Server running on local host http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
