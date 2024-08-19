package controllers

import (
	"database/sql"
	"real-time-forum/backend/modules"
)

// Function for retreive the coms from a user
func AccountComsByAuthors(author string, app *modules.Application) ([]modules.Comment, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT Coms FROM Comments WHERE Author=?", author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []modules.Comment

	for rows.Next() {
		var com modules.Comment
		if err := rows.Scan(&com.Coms); err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
