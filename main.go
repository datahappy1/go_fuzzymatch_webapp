package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize("production")

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	a.Run(":" + port)

}
