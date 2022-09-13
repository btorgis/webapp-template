package handlers

import (
	"time"
	"net/http"
	"html/template"
	"github.com/google/uuid"
	"btorgis.com/webapp/users"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/login.html",
	}
	
	errors := Errors{}
	
	r.ParseForm()
	// Check Login Button Pressed
	if r.Method == "POST" {
		if r.FormValue("submit") == "Login" {
			// Check Login Information
			user.Email = r.FormValue("email")
			user.Password = r.FormValue("password")
			
			// Validate Form Data
			if user.Email == "" {
				errors.UserError = "Email cannot be blank"
			}
			
			if user.Password == "" {
				errors.PasswordError = "Password cannot be blank"
				errors.Email = user.Email
			}

			// Check if User exists
			u := users.User{}
			if users.IsAccountValid(user.Email) {
				// Valid Account, get user
				u = users.GetUser(user.Email)
				// Check Password
				if u.Password == user.Password {
					// Valid Login
					// Create a new random session token
					sessionToken := uuid.NewString()
					expiresAt := time.Now().Add(120 * time.Second)
					
					// Set the token in the session map, along with the session info
					users.Sessions[sessionToken] = users.Session {
						Username: user.Email,
						Expires: expiresAt,
					}
					// Set the client cookie for "session_token"
					http.SetCookie(w, &http.Cookie {
						Name:		"session_token",
						Value:		sessionToken,
						Expires:		expiresAt,
						Path:		"/",
					})
					
					user.IsAuthenticated = true
					http.Redirect(w, r, "/", http.StatusSeeOther)
				} else {
					// Invalid Password
					user.LoginAttempts++
					errors.UserError = "Invalid Password"
				}
				
			} else {
				// Invalid User Credentials
				user.LoginAttempts++
				errors.UserError = "Invalid Email Address"
				/*
				if user.LoginAttempts > 10 {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}*/
			}
		}
	}
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, errors)
	if err != nil {
		panic(err)
	}

}
