package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"real-time-forum/backend/modules"

	"golang.org/x/crypto/bcrypt"
)

// verify the len of password
func VerifyPasswordLenght(Form_password string) bool {
	if len(Form_password) > 8 {
		return true
	} else {
		return false
	}
}

// Checking if the user is in the db
func VerifyUserAlreadyExist(mailInput, usernameInput string, app *modules.Application) bool {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
		defer db.Close()
	}
	defer db.Close()
	var email, username string
	datas, _ := db.Query("SELECT email, Username FROM Auth WHERE email=? OR Username=?", mailInput, usernameInput)
	for datas.Next() {
		err = datas.Scan(&email, &username)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(email, username)
		if email != "" && username != "" {
			return false
		}
	}
	return true
}

// hash the password
func HashPAss(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
