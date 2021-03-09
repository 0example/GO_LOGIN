package main

import (
	"github.com/gorilla/mux" //A powerful HTTP router and URL matcher for building Go web servers
	"net/http" //Package http provides HTTP client and server implementations.
	"../common" //Package common contains various helper functions.
)

var router = mux.NewRouter() //Creating a new router

func main() {
//Registering URL paths and Handlers
	router.HandleFunc("/", common.LoginPageHandler) // GET

	router.HandleFunc("/index", common.IndexPageHandler) // GET
	router.HandleFunc("/login", common.LoginHandler).Methods("POST")

	router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")

	router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")

	http.Handle("/", router) //Handle registers the handler for the given pattern in the DefaultServeMux.
	http.ListenAndServe(":8000", nil)//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
}
