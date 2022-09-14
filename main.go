package main

import (
	"log"
	"flag"
	"time"
	"net/http"

	"btorgis.com/webapp/handlers"
	"btorgis.com/webapp/users"
)


func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	// Initialize Databases
	users.InitDatabase()

	// Setup Route Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/admin", handlers.AdminHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)
	http.Handle("/", mux)

	// This will serve files under http://localhost:8000/static/<filename>
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8000",
		// Enforce Timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	
	log.Fatal(srv.ListenAndServe())
}