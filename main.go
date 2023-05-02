package main

import (
	"github.com/ederson-lehugeur/gin-rest-api/database"
	"github.com/ederson-lehugeur/gin-rest-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
