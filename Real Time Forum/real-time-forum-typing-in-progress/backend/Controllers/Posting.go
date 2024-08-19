package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"real-time-forum/backend/modules"
)

// Displaying all post
func DisplayAllPost(Post []modules.Post, app *modules.Application) ([]modules.Post, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Content")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp modules.Post
		err = rows.Scan(&tmp.ID, &tmp.Categories, &tmp.Title, &tmp.Author, &tmp.Description, &tmp.Post, &tmp.Time)

		if err != nil {
			return nil, err
		}

		tmp.User = app.User

		Post = append(Post, tmp)
	}

	return Post, nil

}

// Displaying post by ids
func DisplayPostbyID(Post modules.Post, app *modules.Application) (modules.Post, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return modules.Post{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Content WHERE ID=?", app.Post.ID)
	if err != nil {
		return modules.Post{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&Post.ID, &Post.Categories, &Post.Title, &Post.Author, &Post.Description, &Post.Post, &Post.Time)
		if err != nil {
			return modules.Post{}, err
		}
	}

	return Post, nil
}

// verify the name of the post
func VerifyPostName(CP modules.CreatePost) bool {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	var title string
	query, _ := db.Query("SELECT Title FROM Content", CP.Title)
	defer query.Close()

	for query.Next() {
		err = query.Scan(&title)
		if err != nil {
			fmt.Println(err)
		}

		if title != "" {

			return false
		}
	}
	return true
}
