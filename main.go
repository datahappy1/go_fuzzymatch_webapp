package main

import (
	"os"
)

func main() {

	go ClearAppRequestData()

	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "production"
	}

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	a := App{}
	a.Initialize(env)
	a.Run(":" + port)

}
