package main

import (
	"server/api"

	"github.com/gorilla/mux"
)

func main() {
	a := api.App{}
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
	a.Run(":8000")
}
