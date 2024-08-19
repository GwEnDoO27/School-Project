package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
)

// Managin the register
func HandleRegister(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var err error

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data modules.Auth
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	if !controllers.VerifyUserAlreadyExist(data.Email, data.Username, app) {
		app.Auth.DisplayErr = true
		app.Auth.Err = "User Already Exists"
		log.Println("User Already Exists")
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
		return
	}

	if !controllers.VerifyPasswordLenght(data.Password) {
		app.Auth.DisplayErr = true
		app.Auth.Err = "Too short Password"
		log.Println("Too short Password")
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
		return
	} else {
		data.Hash, err = controllers.HashPAss(data.Password)
		if err != nil {
			log.Println(err)
		}
		app.PostInf.Logged = true
		app.User.Logged = true
		app.Auth.DisplayErr = false
		controllers.AuthInsertDb(data.Email, data.Username, data.Password, data.Hash, w, app)
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
	}

}
