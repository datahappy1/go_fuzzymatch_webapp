package main

func main() {
	a := App{}
	a.Initialize("root", "", "rest_api_example")

	a.Run(":8080")
	//log.Fatal(http.ListenAndServe(":8080", r))

	//https://github.com/kelvins/GoApiTutorial
	//https://stackoverflow.com/questions/55260250/route-method-becomes-undefined-in-main-package-file/55260736#55260736
	//https://stackoverflow.com/questions/55661916/how-to-import-routes/55663396
}
