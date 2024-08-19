package main

import (
	"fmt"
	"forum/database/controller/categories"
	"forum/database/controller/posts"
	"forum/database/initialisation"
	"forum/handlefunc"
	"forum/structure"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// os.Remove("structure/database.db")

	initialisation.CreateBDD()
	categories.GetCategories()
	posts.GetPosts()

	// Parsing all html files
	structure.Tpl = template.Must(template.New("").ParseGlob("front/**/*.html"))

	// Listen statics files
	fs := http.FileServer(http.Dir("front"))
	http.Handle("/front/", http.StripPrefix("/front", fs))
}

func main() {
	fmt.Print("\033[96m")
	fmt.Println("\033[96mServer started at: http://localhost:8080\033[0m")
	fmt.Print("\033[0m")

	// Create the routes for the differents page
	http.HandleFunc("/", handlefunc.HomeHandle)
	http.HandleFunc("/home", handlefunc.HomeHandle)
	http.HandleFunc("/createPost", handlefunc.CreatePostHandle)
	http.HandleFunc("/profile", handlefunc.ProfileHandle)
	http.HandleFunc("/post/", handlefunc.PostHandle)
	http.HandleFunc("/adminPanel", handlefunc.AdminPanelHandle)
	http.HandleFunc("/disconnect", handlefunc.CompleteDisconnectHandle)
	http.HandleFunc("/bye", handlefunc.DisconnectHandle)

	http.ListenAndServe(":8080", nil)
}
