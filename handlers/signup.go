package handlers

import (

	"time"
	"net/http"
	"html/template"

	"github.com/google/uuid"

	"btorgis.com/webapp/users"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/signup.html",
	}
	
	errors := Errors{}
	
	r.ParseForm()
	// Check Login Button Pressed
	if r.Method == "POST" {
		if r.FormValue("submit") == "Submit" {
			// Check Login Information
			user.Email = r.FormValue("email")
			user.Password = r.FormValue("password")
			
			// Check If User Already Exists
			if users.IsAccountValid(user.Email) {
				errors.UserError = "User already exists"
						
			} else {
				// Add User to Database			
				users.AddUser(user)
				
				// User Credentials Valid
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
				
			}
			
			// Validate Form Data
			if user.Email == "" {
				errors.UserError = "Email cannot be blank"
			}
						
			if user.Password == "" {
				errors.PasswordError = "Password cannot be blank"
				errors.Email = user.Email
			}
			
		
		}
	}
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, errors)
	if err != nil {
		panic(err)
	}

}