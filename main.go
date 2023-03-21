package main

import (
	"github.com/david-luk4s/desafio-dev/adapters/interfaces/api"
	"github.com/david-luk4s/desafio-dev/config/database"
)

func main() {
	// 	//init connection
	database.ConnectionDB()

	//start server
	api.Handler()
}
