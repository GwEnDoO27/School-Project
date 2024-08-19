package handlefunc

import (
	"encoding/json"
	"errors"
	"fmt"
	"forum/cookies"
	"forum/database/controller/users"
	"forum/structure"
	"forum/verificationFunction"
	"io"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// this function manages the registration
func RegisterHandle(w http.ResponseWriter, r *http.Request) error {
	var user structure.Users
	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]

	if r.FormValue("userContent") == "" && r.FormValue("FacebookUserContent") == "" && !strings.HasPrefix(code, "code=") {
		if r.FormValue("psw") != r.FormValue("confirm-psw") {
			return errors.New("password and password confirmation don't match")
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("psw")

		if !verificationFunction.PasswordVerif(password) {
			return errors.New("incorrect password ! the password must contain 8 characters, 1 uppercase letter, 1 special character, 1 number")
		}

		if !verificationFunction.EmailVerif(email) {
			return errors.New("incorrect email format")
		}

		if username == "" || email == "" || password == "" {
			return errors.New("there is an empty field")
		}

		cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
		err := users.AddUser(username, email, string(cryptPassword))
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		user, err = users.GetUser(email, string(cryptPassword))
		if err != nil {
			return err
		}

	} else if r.FormValue("userContent") != "" {
		email, username, _, err := GoogleLogDecoder(r.FormValue("userContent"))
		if err != nil {
			return err
		}

		err = users.AddUser(username, email, "")
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		structure.Html.User = structure.Users{}
		user, err = users.GetUserByEmail(email)
		if err != nil {
			return err
		}
	} else if strings.HasPrefix(code, "code=") {
		datas, err := GetDatasGithub(w, r)
		if err != nil {
			return err
		}

		err = users.AddUser(datas["login"].(string), datas["email"].(string), "")
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		structure.Html.User = structure.Users{}
		user, err = users.GetUserByEmail(datas["email"].(string))
		if err != nil {
			return err
		}
	} else if r.FormValue("FacebookUserContent") != "" {
		email, username, err := FaceBookGetData(r.FormValue("FacebookUserContent"))
		if err != nil {
			return err
		}

		if email == "" {
			return errors.New("your facebook email address isn't in public")
		}

		err = users.AddUser(username, email, "")
		if err != nil {
			return errors.New("username or email is already used by someone")
		}

		structure.Html.User = structure.Users{}
		user, err = users.GetUserByEmail(email)
		if err != nil {
			return err
		}
	}

	users.ConnectUser(structure.Html.User.Id)
	user.Connected = 1
	structure.Html.User = user

	cookies.AddCookies(w)
	return nil
}

func GetDatasGithub(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	code := strings.Split(r.URL.String(), "?")[len(strings.Split(r.URL.String(), "?"))-1]
	codeParam := strings.ReplaceAll(code, "code=", "")

	if code != codeParam {
		client_id := "95d7187086635e790a24"
		client_secret := "5b2562084e53a58c68a95d87045454994c4cc8f1"

		FindAccessToken := "https://github.com/login/oauth/access_token" + "?client_id=" + client_id + "&client_secret=" + client_secret + "&code=" + codeParam
		datas, err := http.Get(FindAccessToken)
		if err != nil {
			return nil, err
		}
		defer datas.Body.Close()

		buffer := make([]byte, 1024)
		_, err = datas.Body.Read(buffer)
		if err != nil {
			return nil, err
		}

		token := (string(buffer[strings.Index(string(buffer), "=")+1 : strings.Index(string(buffer), "&")]))

		req, err := http.NewRequest(
			"GET",
			"https://api.github.com/user",
			nil,
		)
		if err != nil {
			return nil, err
		}
		defer req.Body.Close()

		authorizationHeaderValue := fmt.Sprintf("token %s", token)

		req.Header.Set("Authorization", authorizationHeaderValue)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		scan, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		userData := map[string]interface{}{}

		err = json.Unmarshal(scan, &userData)
		if err != nil {
			return nil, err
		}

		if userData["email"] == nil {
			return nil, errors.New("your github email address is not public")
		}

		return userData, nil
	}

	return nil, errors.New("error while connecting with Github")
}
