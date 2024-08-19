package cookies

import (
	"forum/structure"
	"net/http"
	"time"
)

// Create a cookie for the user
func AddCookies(w http.ResponseWriter) {
	cookieEmail := http.Cookie{
		Name:    "email",
		Value:   structure.Html.User.Email,
		Expires: time.Now().Add(48 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		Path: "/",
	}
	
	http.SetCookie(w, &cookieEmail)
}
