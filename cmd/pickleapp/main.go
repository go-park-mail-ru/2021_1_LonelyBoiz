package main

import (
	"os"
	"server/internal/pickleapp/entryPoint"
)

func main() {

	a := entryPoint.App{}
	config := entryPoint.NewConfig()

	conns, emailConnection := a.InitializeRoutes(config)

	defer func() {
		for _, conn := range conns {
			conn.Close()
		}
		close(emailConnection)
	}()

	err := a.Start()
	if err != nil {
		a.Logger.Error(err)
		os.Exit(1)
	}
}
