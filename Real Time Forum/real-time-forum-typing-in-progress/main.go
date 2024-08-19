package main

import (
	"fmt"
	Database "real-time-forum/backend/Configs"
	"real-time-forum/backend/routes"

	_ "github.com/mattn/go-sqlite3"

	"log"
	"net/http"
	"real-time-forum/backend/modules"
)

func init() {
	Database.CreateDatabase()
}

func main() {
	fmt.Print("\033[37m")
	fmt.Println("Server Started: http://localhost:8080/")
	fmt.Print("\033[0m")

	app := modules.NewApplication()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { routes.DisplayReal(w, r, app) })
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { routes.HandleRegister(w, r, app) })
	http.HandleFunc("/Home", func(w http.ResponseWriter, r *http.Request) { routes.Home(w, r, app) })
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { routes.Handlelogin(w, r, app) })
	http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) { routes.HandleInformations(w, r, app) })
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) { routes.Post(w, r, app) })
	http.HandleFunc("/newpost", func(w http.ResponseWriter, r *http.Request) { routes.CreatePost(w, r, app) })
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { routes.WsHandler(w, r, app) })
	http.HandleFunc("/delete-session", func(w http.ResponseWriter, r *http.Request) { routes.DeleteSessionHandler(w, r, app) })
	http.HandleFunc("/load-message", func(w http.ResponseWriter, r *http.Request) { routes.LoadOldMessage(w, r, app) })
	http.HandleFunc("/all-users", func(w http.ResponseWriter, r *http.Request) { routes.Alluser(w, r, app) })
	http.HandleFunc("/delete-notif", func(w http.ResponseWriter, r *http.Request) { routes.DelNotifFromDb(w, r) })
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) { routes.SendCategories(w, r, app) })
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
