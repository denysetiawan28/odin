package container

import (
	"fmt"
	"odin/src/server/config"
	"odin/src/server/database"
)

type DefaultContainer struct {
	//#register config
	Config *config.DefaultConfig
}

func IntializeContainer() *DefaultContainer {

	config := config.ConfigApps("./resources/")

	//initialize database
	db := database.InitializeDatabase(config.Database)

	row := db.Raw("SELECT 1").Row()
	fmt.Println(row)

	return &DefaultContainer{
		Config: config,
	}
}
