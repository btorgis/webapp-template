package handlers

import (
	"time"
	"net/http"
	"html/template"

	"btorgis.com/webapp/users"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	paths := []string {
		"./templates/index.html",
	}
	user.IsAuthenticated = false
	c, err := r.Cookie("session_token")
	if err != nil {
		// Return unauthorized if Cookie not set
		if err == http.ErrNoCookie {
			//w.WriteHeader(http.StatusUnauthorized)
			//return
		}
		// For any other errors, return a bad request
		w. WriteHeader(http.StatusBadRequest)
		return
	}
	
	// Remove the user session from the session map
	sessionToken := c.Value
	delete(users.Sessions, sessionToken)
	
	// Reset Cookie
	http.SetCookie(w, &http.Cookie {
		Name: "session_token",
		Value: "",
		Expires: time.Now(),
		Path: "/",
	})
	
	t := template.Must(template.ParseFiles(paths...))
	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
	
}