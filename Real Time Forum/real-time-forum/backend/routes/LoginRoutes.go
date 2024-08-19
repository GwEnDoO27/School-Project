package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	Database "real-time-forum/backend/Configs"
	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
)

// managin the login way
func Handlelogin(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	Islogged, err := controllers.VerifyCookie(w, r, app)
	if err != nil {
		fmt.Println(err)
	}

	var data modules.Auth
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	AlreadySession, err := controllers.VerifySession(data.Username, app)
	if err != nil && AlreadySession {
		fmt.Println(fmt.Printf("error when verifying session: %s", err))
	}

	err, Logged := controllers.Login(data.Username, data.Password, app)
	if err != nil {
		log.Println("Login failed:", err)
		app.PostInf.Logged = false
		app.User.Logged = false
		app.Auth.DisplayErr = true
		app.Auth.Err = fmt.Sprint("Login failed : ", err)
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
	} else if !Logged {
		app.PostInf.Logged = false
		app.User.Logged = false
		app.Auth.DisplayErr = true
		app.Auth.Err = "Username or Pasword wrong"
		log.Println("Wrong Connexion token")
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
	} else if Islogged {
		app.PostInf.Logged = false
		app.User.Logged = false
		app.Auth.DisplayErr = true
		app.Auth.Err = "Cookie Already set"
		log.Println("Cookie Already set")
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
	} else if AlreadySession {
		app.PostInf.Logged = false
		app.User.Logged = false
		app.Auth.DisplayErr = true
		app.Auth.Err = "User Already connect"
		log.Println("User Already connect")
	} else {
		controllers.SetingCookie(w, data.Username, app)
		Database.AuthJSOn(data.Username, w, app)
		app.Auth.DisplayErr = false
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(app)
	}
}
