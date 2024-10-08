package handlefunc

import (
	"forum/cookies"
	"forum/database/controller/categories"
	"forum/database/controller/posts"
	"forum/structure"
	"forum/verificationFunction"
	"net/http"
	"strconv"
	"strings"
)

// this function manages the new post creation page
func CreatePostHandle(w http.ResponseWriter, r *http.Request) {
	structure.Html.Error.CreatPostError = ""

	if r.URL.Path != "/createPost" {
		http.Redirect(w, r, "/createPost", http.StatusSeeOther)
		return
	}

	cookies.ConnectCookies(w, r)
	if structure.Html.User.Email == "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
		return
	}
	var Path string
	// Get the image from the page
	ImagePath, handler, err := r.FormFile("Img")
	if err == nil {
		if handler.Size >= 20e+6 {
			structure.Html.Error.CreatPostError = "This size of image is above 20mb"

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}
		Path = posts.FormatImg(ImagePath, handler)
	}

	//call the function for storing the image’
	postTitle := r.FormValue("post-title")
	postDescription := r.FormValue("description")
	postCategories := strings.Join(r.Form["categories"], "|")

	for _, y := range r.Form["categories"] {
		idCat, err := strconv.Atoi(y)
		if err != nil {
			structure.Html.Error.CreatPostError = "one of the categories does not exist"

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}

		str, err := categories.GetCategoriesById(idCat)
		if err != nil || str == "" {
			structure.Html.Error.CreatPostError = "one of the categories does not exist"

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}
	}

	if postCategories != "" && verificationFunction.IsValidMessage(postTitle) && verificationFunction.IsValidMessage(postDescription) {
		err := posts.CreateNewPost(postCategories, structure.Html.User.Id, postTitle, postDescription, Path)
		if err != nil {
			structure.Html.Error.CreatPostError = err.Error()

			structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
			structure.Html.Error.CreatPostError = ""
			return
		}
	} else {
		structure.Html.Error.CreatPostError = "there is an empty field in your post"

		structure.Tpl.ExecuteTemplate(w, "createPost.html", structure.Html)
		structure.Html.Error.CreatPostError = ""
		return
	}

	posts.GetPostsByName(postTitle)

	DisconnectHandle(w, r)
	http.Redirect(w, r, "/post/"+strconv.Itoa(structure.Html.Posts[0].Id), http.StatusSeeOther)
}
