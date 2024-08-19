package controllers

import (
	"log"
	"real-time-forum/backend/modules"
)

// Function for check the solvability of the post
func ValidPost(Title string, Des string, Post string, Cat string, SelectValue string, app *modules.Application) {
	if len(Title) >= 50 || VerifyPostName(app.CreatePost) {
		app.CreatePost.Err = "This Title is already used"
		log.Println("This Title is already used")
		return
	}

	log.Println("Valid title")

	if len(Des) >= 100 {
		app.CreatePost.Err = "Too long Description (> 100 char)"
		log.Println("Too long")
		return
	}

	log.Println("Valid Description")

	if len(Post) >= 1000 {
		app.CreatePost.Err = "This post is too long (> 1000 char)"
		log.Println("Too long post")
		return
	}

	log.Println("Post size accepted")

	if Cat == "" {
		log.Println("added by select")
		return
	} else if !VerifyCatExists(SelectValue) {
		AddCategories(SelectValue)
		log.Printf("%s added by new Category", SelectValue)
		return
	}
}
