package main

func main() {
	a := App{}
	a.Initialize()

	a.Run(":8080")

	//https://github.com/kelvins/GoApiTutorial
	//https://stackoverflow.com/questions/55260250/route-method-becomes-undefined-in-main-package-file/55260736#55260736
	//https://stackoverflow.com/questions/55661916/how-to-import-routes/55663396
}
