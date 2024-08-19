package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
)

// Sending the json post infos to js
func Home(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var posts []modules.Post
	var err error

	posts, err = controllers.DisplayAllPost(posts, app)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to display posts", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
