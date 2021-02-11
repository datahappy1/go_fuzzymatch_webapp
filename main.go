package main

func main() {
	a := App{}
	a.Initialize()
	//a.Initialize("root", "", "rest_api_example")

	a.Run(":8080")

}
