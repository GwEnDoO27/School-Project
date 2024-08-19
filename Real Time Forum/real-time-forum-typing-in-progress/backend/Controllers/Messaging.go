package controllers

import (
	"encoding/json"
	"fmt"
	Database "real-time-forum/backend/Configs"
	"real-time-forum/backend/modules"
	"sort"
	"time"

	"github.com/gorilla/websocket"
)

// SortUserCounts sorts the user counts by the last sent timestamp in descending order
func SortUserCounts(userCounts []modules.UserCount) {
	sort.Slice(userCounts, func(i, j int) bool {
		return userCounts[i].LastSent > userCounts[j].LastSent
	})
}

// Function to broadcast the list of connected users to all clients
func BroadcastConnectedUsers(app *modules.Application) {
	app.Mutex.Lock()
	defer app.Mutex.Unlock()
	for _, session := range app.Sessions {
		if session.Conn != nil {
			var users []string
			for _, s := range app.Sessions {
				if s.Username != session.Username {
					users = append(users, s.Username)
				}
			}
			message := struct {
				Type  string   `json:"type"`
				Users []string `json:"users"`
			}{
				Type:  "connected_users",
				Users: users,
			}
			msg, err := json.Marshal(message)
			if err != nil {
				fmt.Printf("Error marshalling connected users: %v\n", err)
				continue
			}
			if err := session.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("Error writing connected users message: %v\n", err)
			}
		}
	}
}

// Function to broadcast a message to all clients
func BroadcastMessage(app *modules.Application, username string, receivedMessage struct {
	Type    string `json:"type"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}) {

	Database.InsertMessIntoDb(username, receivedMessage.To, receivedMessage.Message, app)
	_, exists := app.Sessions[receivedMessage.To]

	if !exists {
		Database.AddingNotifOnDB(username, receivedMessage.To, app)
	}

	message := struct {
		Type    string `json:"type"`
		From    string `json:"from"`
		To      string `json:"to"`
		Message string `json:"message"`
		Time    string `json:"Time"`
	}{
		Type:    "chat_message",
		From:    username,
		To:      receivedMessage.To,
		Message: receivedMessage.Message,
		Time:    time.Now().Format(time.DateTime),
	}

	msg, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error marshalling chat message: %v\n", err)
		return
	}
	app.Mutex.Lock()
	defer app.Mutex.Unlock()
	for _, session := range app.Sessions {
		if session.Conn != nil {
			if err := session.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("Error writing chat message: %v\n", err)
			}
		}
	}
}

// SelectAllNotifs is the main function to get all notifications
func SelectAllNotifs(user string, app *modules.Application) ([]modules.UserCount, error) {
	db, err := Database.OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	found, err := Database.FetchMessage(db, user)
	if err != nil {
		fmt.Println(err)
	}

	var userCounts []modules.UserCount

	if !found {
		userCounts, err = Database.FetchUsers(db, user)
		if err != nil {
			return nil, err
		}
	} else {
		userCounts, err = Database.FetchUserCounts(db, user)
		if err != nil {
			return nil, err
		}
		existingUsers := make(map[string]bool)
		for _, uc := range userCounts {
			existingUsers[uc.User] = true
		}
		missingUsers, err := Database.FetchMissingUsers(db, user, existingUsers)
		if err != nil {
			return nil, err
		}
		userCounts = append(userCounts, missingUsers...)
		SortUserCounts(userCounts)

	}
	return userCounts, nil
}
