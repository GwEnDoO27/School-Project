package Database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/backend/modules"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Function for creating the db
func CreateDatabase() {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	r := `
    CREATE TABLE IF NOT EXISTS User (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        UserName VARCAHR(20) NOT NULL REFERENCEs "Auth"("username"),
        Email VARCHAR(100) NOT NULL REFERENCES "Auth"("email"),
        Name VARCHAR(20) NOT NULL,
        Password VARCHAR(50) NOT NULL,
        Inscription VARCHAR(10) NOT NULL,
        Birthday VARCHAR(10) NOT NULL,
        Sex VARCHAR(30) NOT NULL,
        Ville VARCHAR(20) NOT NULL,
		Pays VARCHAR(30) NOT NULL,
		Lastname VARCHAR(20) NOT NULL
    );
    CREATE TABLE IF NOT EXISTS AllCategories (
		ID INTEGER PRIMARY KEY,
		Categories TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS Auth (
		ID INTEGER PRIMARY KEY,
		Username  VARCAHR(20) NOT NULL,
		Password VARCAHR(50) NOT NULL,
		email VARCAHR(100) NOT NULL
    );
    CREATE TABLE IF NOT EXISTS Comments (
        ID INTEGER PRIMARY KEY,
		PostID INT NOT NULL REFERENCES "Content"("ID"),
		Coms VARCAHR(200) NOT NULL REFERENCES "Content"("Post"),
		Author VARCAHR(20) NOT NULL,
		Time VARCAHR(30) NOT NULL
    );
	CREATE TABLE IF NOT EXISTS Content (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Categories VARCHAR(100) NOT NULL REFERENCES "AllCategories"("Categories"), 
		Title VARCHAR(100) NOT NULL,
		Author VARCHAR(20) NOT NULL REFERENCES "User"("UserName"),
		Description VARCAHR(300) NOT NULL,
		Post VARCAHR(1000) NOT NULL,
		Time VARCAHR(30) NOT NULL
    );
	CREATE TABLE IF NOT EXISTS messages (
		Username VARCHAR(20) REFERENCES "User"("UserName"),
		Message VARCHAR(100),
		Username2 VARCHAR(20) REFERENCES "User"("UserName"),
		Time TEXT
	);
	CREATE TABLE IF NOT EXISTS notifications (
		sender VARCHAR(20) REFERENCES "Auth"("username"),
		receiver VARCHAR(20) REFERENCES "Auth"("username"),
		Time TEXT
	);
	`

	_, err = db.Exec(r)
	if err != nil {
		log.Println("CREATE ERROR")
		fmt.Println(err)
	}
}

// Insert Message into db
func InsertMessIntoDb(Username string, Username2 string, Mess string, app *modules.Application) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	Time := time.Now().Format(time.DateTime)

	stmt := "INSERT INTO messages (Username,Username2,Message,Time) VALUES (?,?,?,?)"
	_, err = db.Exec(stmt, Username, Username2, Mess, Time)
	if err != nil {
		log.Println(fmt.Printf("%v", err))
	}
}

// Select the message from the db in functionnof users
func SelectMessage(Username string, Username2 string, app *modules.Application) ([]modules.OlderMess, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	query := "SELECT * FROM messages WHERE (Username=? AND Username2=?) OR (Username=? AND Username2=?)  "
	rows, err := db.Query(query, Username, Username2, Username2, Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []modules.OlderMess

	for rows.Next() {
		var message modules.OlderMess
		err = rows.Scan(&message.From, &message.Message, &message.To, &message.Time)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

// Selecting all the users
func SelectAllusers() ([]string, error) {
	var AllUsers []string
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT (UserName) FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)

		if err != nil {
			return nil, err
		}

		AllUsers = append(AllUsers, tmp)

	}

	return AllUsers, nil
}

// Selecting all the message from conversations into two users
func SelectMessageUser(username string) ([]string, error) {
	var MessageUser []string
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT Username, Username2 FROM messages WHERE Username2 = ? OR Username = ?", username, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp modules.Message
		err = rows.Scan(&tmp.Username, &tmp.To)
		if err != nil {
			return nil, err
		}
		if tmp.Username == username {
			MessageUser = append(MessageUser, tmp.To)
		} else if username == tmp.To {
			MessageUser = append(MessageUser, tmp.Username)
		}
	}
	for i := 0; i < len(MessageUser)/2; i++ {
		MessageUser[i], MessageUser[len(MessageUser)-1-i] = MessageUser[len(MessageUser)-1-i], MessageUser[i]
	}

	var FinalTab []string

	for i := range MessageUser {
		if !Istab(FinalTab, MessageUser[i]) {
			FinalTab = append(FinalTab, MessageUser[i])
		}
	}

	return FinalTab, nil
}

// Checking if the user is in the tab
func Istab(tab []string, isintab string) bool {
	for i := range tab {
		if tab[i] == isintab {
			return true
		}
	}
	return false
}

// Order the users in the tab
func OrderUsers(alluser, MessageUser []string) ([]string, error) {
	var finalUser []string

	if MessageUser == nil {
		return alluser, nil
	}
	finalUser = append(finalUser, MessageUser...)

	for _, user := range alluser {
		if !Istab(finalUser, user) {
			finalUser = append(finalUser, user)
		}
	}
	return finalUser, nil
}

// insert Data of new post in db
func GetDataFromNewPost(Title string, Des string, Post string, Cat string, Time string, Name string) error {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := `INSERT INTO Content (Title, Description, Post, Categories, Time, Author) VALUES (?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(stmt, Title, Des, Post, Cat, Time, Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Insert infos in db
func InfosInsertDb(username string, name string, email string, lastname string, birthday string, gender string, town string, country string, time string) error {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := `INSERT INTO User (UserName, Email, Name, lastname, Inscription, Birthday, Sex, Ville, Pays) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(stmt, username, email, name, lastname, time, birthday, gender, town, country)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// Select from the Db the credentials
func AuthJSOn(username string, w http.ResponseWriter, app *modules.Application) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	err = db.QueryRow("SELECT ID,Email,Username FROM User WHERE Username =? OR email=?", username, username).Scan(&app.User.ID, &app.User.Email, &app.User.UserName)
	if err != nil {
		log.Println(err)
	}

	if app.User.ID != 0 {
		app.PostInf.Logged = true
		app.User.Logged = true
	}
}

// Adding the the notif from the db
func AddingNotifOnDB(sender string, receiver string, app *modules.Application) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	stmt := "INSERT INTO notifications (receiver,sender,Time) VALUES (?,?,?)"
	_, err = db.Exec(stmt, receiver, sender, time.Now().Format(time.DateTime))
	if err != nil {
		log.Println(fmt.Printf("%v", err))
	}

}

// Opennign the Db
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", modules.NewConfiguration().DBPath)
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}
	return db, nil
}

// FetchUserCounts retrieves the user counts and the last sent timestamp
func FetchUserCounts(db *sql.DB, user string) ([]modules.UserCount, error) {
	query := `
        SELECT sender, COUNT(*) as count, MAX(Time) as last_sent
        FROM notifications 
        WHERE receiver=?
        GROUP BY sender
        ORDER BY last_sent DESC
    `
	rows, err := db.Query(query, user)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var userCounts []modules.UserCount

	for rows.Next() {
		var sender string
		var count int
		var lastSent string
		err = rows.Scan(&sender, &count, &lastSent)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		///

		userCounts = append(userCounts, modules.UserCount{User: sender, Count: count, LastSent: lastSent})

	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return userCounts, nil
}

// FetchMissingUsers retrieves the users who are not in the first query result and excludes the current user
func FetchMissingUsers(db *sql.DB, user string, existingUsers map[string]bool) ([]modules.UserCount, error) {
	query := `
        SELECT UserName 
        FROM User 
        WHERE UserName NOT IN (
            SELECT sender 
            FROM notifications 
            WHERE receiver=?
        ) AND UserName != ?
    `
	rows, err := db.Query(query, user, user)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var userCounts []modules.UserCount

	for rows.Next() {
		var userName string
		err = rows.Scan(&userName)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		if !existingUsers[userName] {
			userCounts = append(userCounts, modules.UserCount{User: userName, Count: 0, LastSent: ""})
		}
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return userCounts, nil
}

func FetchUsers(db *sql.DB, user string) ([]modules.UserCount, error) {
	query := `
        SELECT UserName 
        FROM User 
		ORDER BY UserName 
    `
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var userCounts []modules.UserCount

	for rows.Next() {
		var userName modules.UserCount
		err = rows.Scan(&userName.User)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		userCounts = append(userCounts, userName)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return userCounts, nil
}

func FetchMessage(db *sql.DB, user string) (bool, error) {
	query := `
        SELECT Username, Username2
        FROM messages 
        WHERE Username=? 
    `
	rows, err := db.Query(query, user)
	if err != nil {
		log.Println("Error executing query:", err)
		return false, err
	}
	defer rows.Close()

	var userName, userName2 string
	found := false

	for rows.Next() {
		err = rows.Scan(&userName, &userName2)
		if err != nil {
			log.Println("Error scanning row:", err)
			return false, err
		}
		// If a row is found, set found to true
		found = true
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return false, err
	}

	return found, nil
}
