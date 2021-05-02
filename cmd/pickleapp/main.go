package main

import (
	"math/rand"
	"net/http"
	"os"
	"server/internal/pickleapp/entryPoint"
)

func main() {
	a := entryPoint.App{}
	config := entryPoint.NewConfig()

	conn := a.InitializeRoutes(config)

	defer conn.Close()

	err := a.Start()
	if err != nil {
		a.Logger.Error(err)
		os.Exit(1)
	}
}
