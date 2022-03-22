package main

import (
	"github.com/hoanggggg5/fiber-postgre/config"
	"github.com/hoanggggg5/fiber-postgre/models"
	"github.com/hoanggggg5/fiber-postgre/routes"
)

func main() {
	config.InitConfig()

	config.Database.AutoMigrate(models.User{}, models.Music{})

	routes.InitRouter()
}