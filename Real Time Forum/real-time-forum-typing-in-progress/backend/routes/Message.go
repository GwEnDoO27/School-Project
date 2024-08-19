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
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Handler to delete the session
func DeleteSessionHandler(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	_, username, err := controllers.ParseCookie(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	app.Mutex.Lock()
	delete(app.Sessions, username)
	app.Mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

// WebSocket handler
func WsHandler(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	sessionToken, username, err := controllers.ParseCookie(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	app.Mutex.Lock()
	session, exists := app.Sessions[username]
	app.Mutex.Unlock()
	if !exists || session.Username != username || session.Expire.Before(time.Now()) {
		http.Error(w, "Session expired", http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer func() {
		conn.Close()
		app.Mutex.Lock()
		delete(app.Sessions, sessionToken)
		app.Mutex.Unlock()
		controllers.BroadcastConnectedUsers(app)
	}()
	app.Mutex.Lock()
	session.Conn = conn
	app.Sessions[sessionToken] = session
	app.Mutex.Unlock()
	controllers.BroadcastConnectedUsers(app)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}
		var receivedMessage struct {
			Type    string `json:"type"`
			From    string `json:"from"`
			To      string `json:"to"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(msg, &receivedMessage); err != nil {
			fmt.Printf("Error unmarshalling message: %v\n", err)
			continue
		}

		if receivedMessage.Type != "request_connected_users" && receivedMessage.Type != "typing" {
			controllers.BroadcastMessage(app, username, receivedMessage)
		} else if receivedMessage.Type == "typing" {
			Istyping(receivedMessage.To, receivedMessage.From, app)
		}
	}
}

func LoadOldMessage(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var err error
	var OldMessages []modules.OlderMess
	_, username, err := controllers.ParseCookie(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	var data modules.OlderMess
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	OldMessages, err = Database.SelectMessage(username, data.To, app)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(OldMessages)
}

func Alluser(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var err error
	var Allusers []modules.UserCount

	_, username, err := controllers.ParseCookie(r)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	Allusers, err = controllers.SelectAllNotifs(username, app)
	if err != nil {
		fmt.Println("Error in allusers loop :", err)
	}
	json.NewEncoder(w).Encode(Allusers)
}

func Istyping(receiver, from string, app *modules.Application) {
	//session, exists := app.Sessions[receiver]

	for _, i := range app.Sessions {
		if i.Username != from {
			continue
		}

		fmt.Println("i", i)

		// Check if the connection is nil
		if i.Conn == nil {
			log.Printf("Error: WebSocket connection for user %s is nil", receiver)
			continue
		}

		app.Mutex.Lock()
		defer app.Mutex.Unlock()
		if err := i.Conn.WriteJSON(map[string]string{"from": from, "type": "typing"}); err != nil {
			log.Printf("Error writing JSON to %s: %v", receiver, err)
		}
	}

}
