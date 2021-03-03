package main

import (
	"os"
	"server/api"
)

func main() {
	a := api.App{}
	config := api.NewConfig()
	a.InitializeRoutes(config)
	err := a.Start()
	if err != nil {
		os.Exit(1)
	}
}
