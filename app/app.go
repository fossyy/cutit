package app

import (
	"github.com/fossyy/cutit/db"
	"net/http"
)

type App struct {
	http.Server
	Database db.Database
}

var Server *App

func NewApp(address string, handler http.Handler, database db.Database) *App {
	return &App{
		Server: http.Server{
			Addr:    address,
			Handler: handler,
		},
		Database: database,
	}
}
