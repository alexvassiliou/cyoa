package main

import (
	"net/http"
)

func main() {
	// filePath := "gopher.json"

	// adventures := Parse(filePath)

	// for k := range adventures {
	// 	fmt.Println(k)
	// }

	var testHandler Adventure

	http.Handle("/test", testHandler)

	err := http.ListenAndServe(":8080", nil)
	CheckError(err)

}
