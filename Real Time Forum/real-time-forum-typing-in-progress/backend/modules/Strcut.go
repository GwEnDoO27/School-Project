package modules

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Configuration holds application-wide settings.
type Configuration struct {
	DBPath string
}

// NewConfiguration creates a new configuration instance.
func NewConfiguration() *Configuration {
	return &Configuration{
		DBPath: "./backend/Configs/database.sqlite",
	}
}

// Auth struct
type Auth struct {
	ID         int    `json:"ID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Hash       string `json:"hash"`
	Email      string `json:"email"`
	DisplayErr bool   `json:"DisErr"`
	Err        string `json:"Err"`
	Logged     bool   `json:"Logged"`
}

// User struct
type User struct {
	ID          int    `json:"ID"`
	UserName    string `json:"UserName"`
	Lastname    string `json:"Lastname"`
	Password    string `json:"Password"`
	Hash        string
	Email       string `json:"Email"`
	Name        string `json:"name"`
	Inscription string `json:"Inscription"`
	Birthday    string `json:"Birthday"`
	Sex         string `json:"Gender"`
	Ville       string `json:"Town"`
	Pays        string `json:"Country"`
	Logged      bool   `json:"Logged"`
}

// CreatePost struct
type CreatePost struct {
	ID          int    `json:"ID"`
	Post        string `json:"Post"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Categories  string `json:"Categories"`
	Author      string `json:"Author"`
	Time        string `json:"Time"`
	Img         string `json:"Img"`
	SelectValue string `json:"SelectValue"`
	Err         string `json:"Err"`
}

// struct with all categories
type AllCategories struct {
	ID         int
	Categories string
}

// post struct
type Post struct {
	ID          int    `json:"ID"`
	Categories  string `json:"Categories"`
	Title       string `json:"Title"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
	Post        string `json:"Post"`
	Comments    string `json:"Comments"`
	Time        string `json:"Time"`
	User        User   `json:"User"`
}

// Comme struct
type Comment struct {
	ID     int    `json:"ID"`
	Coms   string `json:"Coms"`
	Author string `json:"Author"`
	PostID int    `json:"PostId"`
	Time   string `json:"Time"`
}

// Strcuct for post and coms
type PostInf struct {
	Post    Post
	Comment []Comment
	Logged  bool
}

type UserAccount struct {
	Post    []Post
	Comment []Comment
	User    User
}

// struct for manage the connected users
type Session struct {
	IdSession string
	Username  string `json:"username"`
	Expire    time.Time
	Conn      *websocket.Conn
	Connected bool
}

// Cookie struct
type Cookie struct {
	Name  string
	Value string
}

// Message struct
type Message struct {
	Type         string    `json:"type"`
	Username     string    `json:"username"`
	To           string    `json:"to"`
	Messsage     string    `json:"message"`
	Users        []string  `json:"users,omitempty"`
	LastMessages OlderMess `json:"LastMessages"`
	Time         string    `json:"Time"`
}

// OlderMess struct
type OlderMess struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"Oldmessages"`
	Time    string `json:"Time"`
}

// Notification struct
type Notification struct {
	Sender   string `json:"from"`
	Receiver string `json:"to"`
}

// Strcut with al users and their notifications
type Alluser struct {
	Name  string `json:"name"`
	Notif int    `json:"notif"`
}
// Application holds all application-wide variables and configurations.
type Application struct {
	Config       *Configuration
	Auth         Auth
	User         User
	CreatePost   CreatePost
	AllCat       AllCategories
	Post         Post
	Comment      Comment
	Message      Message
	PostInf      PostInf
	Session      Session
	Mutex        sync.Mutex
	OlderMess    OlderMess
	Notification Notification
	Sessions     map[string]Session
}

// struct with user, notif and last send message
type UserCount struct {
	User     string `json:"user"`
	Count    int    `json:"notif"`
	LastSent string
}

// NewApplication creates a new instance of Application with default values.
func NewApplication() *Application {
	return &Application{
		Config:   NewConfiguration(),
		Sessions: make(map[string]Session),
	}
}
