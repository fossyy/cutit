package main

import (
	"fmt"
	"github.com/fossyy/cutit/app"
	"github.com/fossyy/cutit/db"
	_ "github.com/fossyy/cutit/db"
	"github.com/fossyy/cutit/routes"
)

func main() {
	handler := routes.Setup()
	database := db.NewMYSQLdb("root", "Password123", "localhost", "3306", "testaja")
	server := app.NewApp("localhost:8000", handler, database)

	app.Server = server

	fmt.Printf("Listening on http://%s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
