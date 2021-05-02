package main

import (
	"os"
	"server/internal/pickleapp/entryPoint"
)

func main() {
	a := entryPoint.App{}
	config := entryPoint.NewConfig()

	conns := a.InitializeRoutes(config)

	defer func() {
		for _, conn := range conns {
			conn.Close()
		}
	}()

	err := a.Start()
	if err != nil {
		a.Logger.Error(err)
		os.Exit(1)
	}
}
