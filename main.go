package main

import (
	"os"
)

func main() {

	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "production"
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	a := App{}

	go a.ClearAppRequestData()
	a.Initialize(env)
	a.Run(":" + port)

}
