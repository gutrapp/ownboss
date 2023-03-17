package main

import (
	"ownboss/models"
	"ownboss/routes"
)

func main() {
	models.Setup()
	routes.SetupAndListen()
}
