package handlers

import (

	"net/http"
	"html/template"

	"btorgis.com/webapp/users"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/signup.html",
	}
	
	r.ParseForm()
	// Check Login Button Pressed
	if r.Method == "POST" {
		if r.FormValue("submit") == "Submit" {
			// Check Login Information
			user.Credentials.Username = r.FormValue("email")
			user.Credentials.Password = r.FormValue("password")
		
			// Check Valid Input
			
			// Check If User Already Exists
			
			// Add User to Database
			users.UserDB[user.Credentials.Username] = user.Credentials.Password
			
		
		}
	}
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}