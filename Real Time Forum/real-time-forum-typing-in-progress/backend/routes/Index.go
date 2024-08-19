package routes

import (
	"net/http"
	"real-time-forum/backend/modules"
	"real-time-forum/frontend/Templates"
)

// Execute the template
func DisplayReal(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var path = "frontend/views/index.html"
	Templates.Rendertemp(w, path)
}
