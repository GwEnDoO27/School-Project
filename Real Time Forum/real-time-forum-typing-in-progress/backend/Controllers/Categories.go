package controllers

import (
	"database/sql"
	"log"
	"real-time-forum/backend/modules"
)

// Cheking if the category is already in the db
func VerifyCatExists(category string) bool {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	var Categories string
	row := db.QueryRow("SELECT Categories FROM AllCategories WHERE Categories=?", category)
	err = row.Scan(&Categories)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return false
	}
	return true
}

// Adding a category in the db
func AddCategories(category string) bool {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	insert := "INSERT INTO AllCategories (Categories) VALUES (?)"
	_, err = db.Exec(insert, category)
	if err != nil {
		log.Println(err)
		log.Println("This category already exists")
		return false
	}
	return true
}

// Displaying all the categories
func DisplayCategories(All []modules.AllCategories) ([]modules.AllCategories, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM AllCategories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp modules.AllCategories
		err = rows.Scan(&tmp.ID, &tmp.Categories)

		if err != nil {
			return nil, err
		}

		All = append(All, tmp)
	}

	return All, nil
}
