package main

import (
	"log"
	"flag"
	"time"
	"net/http"
	
	//"github.com/gorilla/mux"

	"btorgis.com/webapp/handlers"
)


func main() {
	var dir string
	
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	//r := mux.NewRouter()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/logout", handlers.LogoutHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)
	http.Handle("/", mux)

	// This will serve files under http://localhost:8000/static/<filename>
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	
	log.Fatal(srv.ListenAndServe())
}