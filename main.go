package main

import (
	"github.com/holanda-lucas/gotasks-back/database"
	"github.com/holanda-lucas/gotasks-back/routes"
	"github.com/holanda-lucas/gotasks-back/util"
)

func main() {
	database.ConnectWithDatabase()

	util.StartTokenDisposer()
	
	routes.LoadRoutes()
}