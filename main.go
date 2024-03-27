package main

import (
	"fmt"
	_ "github.com/fossyy/cutit/db"
	"github.com/fossyy/cutit/handler/alias"
	errorHandler "github.com/fossyy/cutit/handler/error"
	indexHandler "github.com/fossyy/cutit/handler/index"
	logoutHandler "github.com/fossyy/cutit/handler/logout"
	miscHandler "github.com/fossyy/cutit/handler/misc"
	signinHandler "github.com/fossyy/cutit/handler/signin"
	signupHandler "github.com/fossyy/cutit/handler/signup"
	userHandler "github.com/fossyy/cutit/handler/user"
	"github.com/fossyy/cutit/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	serverAddr := "localhost:8000"
	handler := mux.NewRouter()
	server := http.Server{
		Addr:    serverAddr,
		Handler: middleware.Handler(handler),
	}

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Auth(indexHandler.GET, w, r)
		case http.MethodPost:
			middleware.Auth(indexHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Guest(signinHandler.GET, w, r)
		case http.MethodPost:
			middleware.Guest(signinHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Guest(signupHandler.GET, w, r)
		case http.MethodPost:
			middleware.Guest(signupHandler.POST, w, r)
		}
	})

	handler.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Auth(userHandler.GET, w, r)
		}
	})

	handler.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		middleware.Auth(logoutHandler.GET, w, r)
	})

	handler.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		miscHandler.Robot(w, r)
	})

	handler.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		miscHandler.Favicon(w, r)
	})

	handler.HandleFunc("/{alias}", func(w http.ResponseWriter, r *http.Request) {
		alias.ALL(w, r)
	})

	handler.NotFoundHandler = http.HandlerFunc(errorHandler.ALL)

	fileServer := http.FileServer(http.Dir("./public"))
	handler.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	fmt.Printf("Listening on http://%s\n", serverAddr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
