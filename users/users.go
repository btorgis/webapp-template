package users

import (
	"fmt"
	"time"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"btorgis.com/webapp/data"
)

// Stores the user sessions. Use Database for larger scale app
var Sessions = map[string]Session{}

type User struct {
	gorm.Model
	Email				string	`json:"email"`
	Password			string	`json:"password"`
	Role				string	`json:"role"`
	IsAuthenticated  	bool	`json:"authenticated"`
	LoginAttempts		int		`json:"logins"`
}

func InitDatabase() {
	var err error
	// Open datbase connection
	data.UserDB, err = gorm.Open("sqlite3", "./data/users.db")
	if err != nil {
		panic("Failed to connect to User Database")
	}
	
	fmt.Println("User Database Opened Successfully")
	
	data.UserDB.AutoMigrate(&User{})
	fmt.Println("Database Migrated: User")
}

func GetUser (email string) User {
	var user User
	data.UserDB.Where("email = ?", email).First(&user)
	return user
}

func GetUsers (email string) []User {
	var users []User
	data.UserDB.Where("email LIKE ?", "%%" + email + "%%").Find(&users)
	return users
}

// Add User to Database
func AddUser (user User) error {
	data.UserDB.Create(&user)
	return nil
}

// Delete User from Database
func DeleteUser (user User) error {
	
	data.UserDB.Delete(&user)
	return nil
}

// Check if User Exists
func IsAccountValid (email string) bool {
	var user User
	data.UserDB.Where("email = ?", email).First(&user)
	
	if user.Email == email {
		return true
	} else {
		return false
	}
}

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
