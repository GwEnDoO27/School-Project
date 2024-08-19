package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
)

// Sending the json of categories for the js
func SendCategories(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var Cats []modules.AllCategories
	var err error

	Cats, err = controllers.DisplayCategories(Cats)
	if err != nil {
		fmt.Println("Error when find categories", err)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Cats)
}
