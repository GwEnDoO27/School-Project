package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"real-time-forum/backend/modules"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Function for setting the coockies
func SetingCookie(w http.ResponseWriter, Username string, app *modules.Application) {
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Minute)
	app.Sessions[Username] = modules.Session{
		IdSession: sessionToken,
		Username:  Username,
		Expire:    expiresAt,
		Connected: true,
	}

	// Create the cookie value combining the session token and username
	cookieValue := url.QueryEscape(sessionToken + ":" + Username)

	http.SetCookie(w, &http.Cookie{
		Name:     "session-token",
		Value:    cookieValue,
		Expires:  expiresAt,
		HttpOnly: false, // Prevents JavaScript access
		Secure:   false, // Use Secure in production with HTTPS
		Path:     "/",   // Define the cookie path
	})
}

// Fucntion for parse the cookie in retriev his content
func ParseCookie(r *http.Request) (sessionToken string, username string, err error) {
	cookie, err := r.Cookie("session-token")
	if err != nil {
		return "", "", err
	}

	decodedValue, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return "", "", err
	}

	parts := strings.SplitN(decodedValue, ":", 2)
	if len(parts) != 2 {
		return "", "", errors.New("invalid cookie format")
	}

	return parts[0], parts[1], nil
}

// Verifying if the cookie is already set
func VerifyCookie(w http.ResponseWriter, r *http.Request, app *modules.Application) (bool, error) {
	c, err := r.Cookie(app.Auth.Username)
	if err != nil {
		if err == http.ErrNoCookie {
			return false, err
		}
	}

	sessionToken := c.Value

	userSession, exists := app.Sessions[sessionToken]
	if !exists {

		return false, err

	}
	app.Post.Author = userSession.Username
	return true, err
}
