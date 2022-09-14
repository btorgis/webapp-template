package handlers

import (
	
	"fmt"
	"net/http"
	"html/template"

	"btorgis.com/webapp/users"
)

type TemplateData struct {
	User				users.User
	Users				[]users.User
	IsAuthenticated		bool
	Role				string
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/admin.html",
	}
	user.IsAuthenticated = false
	
	td := TemplateData{}
	var query string
	
	// Process Form Data
	r.ParseForm()
	if r.Method == "POST" {
		// Get Users from Database
		query = r.FormValue("search")
		us := users.GetUsers(query)
		td.Users = us
	}
	
	if r.Method == "GET" {
		// Delete User from Database
		if r.URL.Query().Get("delete") == "true" {
			email := r.URL.Query().Get("email")
			u := users.GetUser(email)
			if u.Email != "" {
				users.DeleteUser(u)
			}

			// Refresh the query results after deletion
			//query := r.FormValue("search")
			us := users.GetUsers(query)
			td.Users = us
			
		}

	}
	
	for key, value := range users.Sessions {
		fmt.Println("Sessions: ", key, value)
	}
	

	if cookie, err := r.Cookie("session_token"); err == nil {
		fmt.Println("token: %v", cookie)
		// Valid Cookie
		// Check Valid Session TODO		
		td.Role = "admin"
		td.IsAuthenticated = true
	} else {
		td.Role = "nobody"
	}
	

	
	
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, td)
	if err != nil {
		panic(err)
	}
}