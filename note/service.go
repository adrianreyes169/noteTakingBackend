package note

import (
	"database/sql"
	"errors"
	"time"
)

func CreateNoteService(db *sql.DB, n Note) (int64, error) {
	if n.Title == "" {
		return 0, errors.New("title cannot be empty")
	}
	if len(n.Title) > 100 {
		return 0, errors.New("title is too long")
	}
	if n.Content == "" {
		return 0, errors.New("content cannot be empty")
	}
	if len(n.Content) > 1000 {
		return 0, errors.New("too much content")
	}

	now := time.Now().UTC()
	n.CreatedAt = now
	n.UpdatedAt = now

	ID, err := CreateNote(db, n)

	if err != nil {
		return 0, err
	}

	return ID, nil
}

func ShowNotesService(db *sql.DB) ([]Note, error) {
	note, err := ShowNotes(db)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func ShowNotesByIDService(db *sql.DB, ID int) (Note, error) {
	if ID <= 0 {
		return Note{}, errors.New("ID is invalid")
	}

	note, err := ShowNoteByID(db, ID)

	if err != nil {
		return Note{}, err
	}

	return note, nil
}
