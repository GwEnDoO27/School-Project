package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	Database "real-time-forum/backend/Configs"
	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
)

// Getting the new post data from js
func CreatePost(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	_, username, err := controllers.ParseCookie(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	var data modules.CreatePost
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	json.Unmarshal(body, &data)
	controllers.ValidPost(data.Title, data.Description, data.Post, data.Categories, data.SelectValue, app)
	if data.Categories == "" {
		data.Categories = data.SelectValue
	}
	err = Database.GetDataFromNewPost(data.Title, data.Description, data.Post, data.Categories, data.Time, username)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(app)

}
