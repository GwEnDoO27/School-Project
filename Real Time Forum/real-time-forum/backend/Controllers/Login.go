package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/backend/modules"

	"golang.org/x/crypto/bcrypt"
)

// Fucntion who check the credentials
func Login(Form_username string, Form_password string, app *modules.Application) (error, bool) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err), false
	}
	defer db.Close()

	err = db.QueryRow("SELECT Hash FROM Auth WHERE UserName=? OR Email=?", Form_username, Form_username).Scan(&app.Auth.Hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found"), false
		}
		return fmt.Errorf("failed to query database: %w", err), false
	}

	// Compare hashed password
	if !CheckPass(Form_password, app.Auth.Hash) {
		return fmt.Errorf("incorrect password"), false
	}

	return nil, true
}

// Comparing hash passwords
func CheckPass(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// verify if the user is in the connected user map
func VerifySession(username string, app *modules.Application) (bool, error) {
	_, exists := app.Sessions[username]
	if exists {
		log.Printf("User %s is already connected", username)
		return true, nil
	}
	return false, nil
}

// Insert into db Auth informations
func AuthInsertDb(email string, name string, pass string, hash string, w http.ResponseWriter, app *modules.Application) {

	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	stmt := "INSERT INTO Auth (email,Password,Username,Hash) VALUES (?,?,?,?)"
	_, err = db.Exec(stmt, email, pass, name, hash)
	if err != nil {
		log.Println(fmt.Printf("%v", err))
	}
	SetingCookie(w, name, app)
	app.PostInf.Logged = true
	app.Auth.Email = email
	app.Auth.Username = name
}
