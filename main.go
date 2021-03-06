package main

import (
	"fmt"
	"os"
	"server/api"
)

func main() {
	a := api.App{}
	config := api.NewConfig()
	a.InitializeRoutes(config)
	err := a.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
