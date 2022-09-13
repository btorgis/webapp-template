package handlers

import (
	"net/http"
	"html/template"

	"btorgis.com/webapp/users"
)

type Errors struct {
	UserError			string
	PasswordError		string
	Email				string
	Password			string
}

var user = users.User{IsAuthenticated: false}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/index.html",
	}
	
	t := template.Must(template.ParseFiles(paths...))
	err := t.Execute(w, user)
	if err != nil {
		panic(err)
	}
}