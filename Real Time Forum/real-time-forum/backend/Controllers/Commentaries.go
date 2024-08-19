package controllers

import (
	"database/sql"
	"log"
	"real-time-forum/backend/modules"
	"strings"
)

// Verifying com format
func VerifyComment(comments string) bool {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return false
	}
	defer db.Close()
	if strings.HasPrefix(comments, " ") {
		return false
	}
	if len(comments) > 1000 {
		return false
	}
	return true

}

// Fucntion for displaying all the coms
func DisplayAllComments(Id int) ([]modules.Comment, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT ID,Coms,Author FROM Comments WHERE PostID=?", Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Post []modules.Comment

	for rows.Next() {
		var tmp modules.Comment
		err = rows.Scan(&tmp.ID, &tmp.Coms, &tmp.Author)

		if err != nil {
			return nil, err
		}
		Post = append(Post, tmp)
	}

	return Post, nil
}

// Function for sending the coms
func SendCom(coms string, Autor string, Idpost int, time string, app *modules.Application) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	if Autor == "" {
		Autor = "Anonymous"
		goto send
	} else {
		goto send
	}
send:
	stmt := "INSERT INTO Comments (Coms,Author,PostID,Time) VALUES(?,?,?,?)"
	_, err = db.Exec(stmt, coms, Autor, Idpost, time)
	if err != nil {
		log.Println(err)
	}
	app.Comment.Author = ""
}
