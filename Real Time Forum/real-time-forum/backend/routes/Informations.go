package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	Database "real-time-forum/backend/Configs"
	"real-time-forum/backend/modules"
	"time"
)

// Getting the infomation data from the js
func HandleInformations(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data modules.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	app.User.Name = data.Name
	app.User.Lastname = data.Lastname
	app.User.UserName = app.Auth.Username
	app.User.Birthday = data.Birthday
	app.User.Sex = data.Sex
	app.User.Ville = data.Ville
	app.User.Pays = data.Pays
	app.User.Inscription = time.Now().Format(time.ANSIC)

	Database.InfosInsertDb(app.Auth.Username, app.Auth.Email, app.User.Name, app.User.Lastname, app.User.Birthday, app.User.Sex, app.User.Ville, app.User.Pays, app.User.Inscription)
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(app.PostInf.Logged)
}
