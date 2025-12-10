package note

import (
	"database/sql"
)

func CreateNote(db *sql.DB, n Note) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO notes (title,content,createdAt,updatedAt) VALUES (?,?,?,?)",
		n.Title,
		n.Content,
		n.CreatedAt,
		n.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func ShowNotes(db *sql.DB) ([]Note, error) {
	var notes []Note
	rows, err := db.Query("SELECT * FROM notes")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n Note
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func ShowNoteByID(db *sql.DB, ID int) (Note, error) {
	var n Note
	row := db.QueryRow("SELECT * FROM notes WHERE id = (?)",
		ID)

	err2 := row.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)

	if err2 != nil {
		return n, err2
	}

	return n, nil
}
