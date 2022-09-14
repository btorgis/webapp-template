package users

import (
	//"fmt"
	"time"
	
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	//"btorgis.com/webapp/data"
)

// Stores the user sessions. Use Database for larger scale app
var Sessions = map[string]Session{}


// Each session contains the username and expiration
type Session struct {
	Username		string
	Expires 		time.Time
}

// Method to determine if the session has expired
func (s Session) isExpired() bool {
	return s.Expires.Before(time.Now())
}


/*
// Method to determine if session is valid
func (s Session) isValid() bool {

}
*/
