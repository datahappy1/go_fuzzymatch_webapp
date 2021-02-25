package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize("production")
	//a.Initialize("root", "", "rest_api_example")

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	a.Run(":" + port)

}
