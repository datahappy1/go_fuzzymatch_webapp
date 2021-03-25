package main

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api"
	"os"
)

func main() {

	env, ok := os.LookupEnv("environment")
	if !ok {
		env = "production"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	a := api.App{}

	a.InitializeAPI(env)
	a.InitializeStatic()
	a.InitializeDB()

	go a.ClearAppRequestData()

	a.Run(":" + port)

}
