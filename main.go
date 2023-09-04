package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"notesTask/database"
	_ "notesTask/docs"
	routes "notesTask/routing"
)

// @title 	Notes Task
// @version	1.0
// @description Main entry point for api

// @host 	localhost:8080
// @BasePath /
func main() {

	database.InitDB()

	serverPort := viper.GetString("server.port")

	r := routes.SetupRouter()

	fmt.Printf("Server is running on :%s\n", serverPort)
	if err := r.Run(fmt.Sprintf(":%s", serverPort)); err != nil {
		log.Fatal(err)
	}
}
