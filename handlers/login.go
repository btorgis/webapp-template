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
	
	r.ParseForm()
	// Check Login Button Pressed
	if r.Method == "POST" {
		if r.FormValue("submit") == "Login" {
			// Check Login Information
			user.Credentials.Username = r.FormValue("email")
			user.Credentials.Password = r.FormValue("password")
			
			// Get the password from map
			expectedPassword, ok := users.UserDB[user.Credentials.Username]
			// Check Password
			if !ok || expectedPassword != user.Credentials.Password {
				// Invalid User Credentials
				user.LoginAttempts++
				if user.LoginAttempts > 5 {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			} else {
				// User Credentials Valid
				// Create a new random session token
				sessionToken := uuid.NewString()
				expiresAt := time.Now().Add(120 * time.Second)
				
				// Set the token in the session map, along with the session info
				users.Sessions[sessionToken] = users.Session {
					Username: user.Credentials.Username,
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
			}
		}
	}
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}
