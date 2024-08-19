package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"real-time-forum/backend/modules"
)

// Deleting the the notif from the db
func DelNotifFromDb(w http.ResponseWriter, r *http.Request) {
	var data modules.Notification
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(body, &data)

	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	stmt := "DELETE FROM notifications WHERE receiver=? AND sender=?"
	_, err = db.Exec(stmt, data.Sender, data.Receiver)
	if err != nil {
		log.Println(fmt.Printf("%v", err))
	}
}
