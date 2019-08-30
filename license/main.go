package main

import (
	routers "license/api/routers"
	"license/config"
	"license/storage"
)

func main() {
	configuration, _ := config.InitConfiguration()

	config.InitConfig()
	//Create Datbase
	storage.InitDB(configuration)
	r := routers.SetupRouter()
	// running
	r.Run(":8085")
}
