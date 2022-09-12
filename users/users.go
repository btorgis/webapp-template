package users

import (
	"time"
)

var UserDB = map[string]string {
	"btorgis@gmail.com": "123",
	"btorgis": "123",
}

// Stores the user sessions. Use Database for larger scale app
var Sessions = map[string]Session{}

// Each session contains the username and expiration
type Session struct {
	Username		string
	Expires 		time.Time
}

// Stuct Model of the user in the request body
type Credentials struct {
	Password	string	`json:"password"`
	Username	string	`json:"username"`
}

type User struct {
	IsAuthenticated  	bool
	LoginAttempts		int
	Credentials			Credentials
}

// Method to determine if the session has expired
func (s Session) isExpired() bool {
	return s.Expires.Before(time.Now())
}
