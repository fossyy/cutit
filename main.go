package main

import (
	"fmt"
	"github.com/fossyy/cutit/app"
	"github.com/fossyy/cutit/db"
	_ "github.com/fossyy/cutit/db"
	"github.com/fossyy/cutit/routes"
	"github.com/fossyy/cutit/utils"
)

func main() {
	handler := routes.Setup()
	serverAddr := fmt.Sprintf("%s:%s", utils.Getenv("SERVER_HOST"), utils.Getenv("SERVER_PORT"))

	dbUser := utils.Getenv("DB_USERNAME")
	dbPass := utils.Getenv("DB_PASSWORD")
	dbHost := utils.Getenv("DB_HOST")
	dbPort := utils.Getenv("DB_PORT")
	dbName := utils.Getenv("DB_NAME")
	database := db.NewMYSQLdb(dbUser, dbPass, dbHost, dbPort, dbName)
	server := app.NewApp(serverAddr, handler, database)

	app.Server = server

	fmt.Printf("Listening on http://%s\n", serverAddr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
