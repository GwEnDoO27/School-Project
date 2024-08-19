package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	controllers "real-time-forum/backend/Controllers"
	"real-time-forum/backend/modules"
	"strconv"
	"time"
)

// Getting the post info and sending post data for the js
func Post(w http.ResponseWriter, r *http.Request, app *modules.Application) {
	var err error

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//var data modules.PostInf
	var datas = make(map[string]string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &datas)
	if err != nil {
		log.Println(err)
	}

	app.Post.ID, _ = strconv.Atoi(datas["id"])

	app.Comment.PostID = app.Post.ID

	app.PostInf.Post, err = controllers.DisplayPostbyID(app.Post, app)
	if err != nil {
		log.Println(err)
	}
	app.PostInf.Comment = append(app.PostInf.Comment, app.Comment)
	app.Comment.Coms = datas["Coms"]
	//fmt.Println(fmt.Printf(" Commentaires: %s", app..Coms))
	app.Comment.Author = app.Auth.Username
	app.Comment.Time = time.Now().Format(time.ANSIC)
	_, isok := datas["Posting"]
	if isok {
		controllers.SendCom(app.Comment.Coms, app.Comment.Author, app.Comment.PostID, app.Comment.Time, app)
	}
	var coms []modules.Comment
	coms, err = controllers.DisplayAllComments(app.PostInf.Post.ID)
	if err != nil {
		fmt.Println(err)
	}
	app.PostInf.Comment = coms
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(app.PostInf)
}
