package main

import (
	"fmt"
	"log"
	"net/http"
)

var a Adventure

func main() {
	filePath := "gopher.json"

	story := Parse(filePath)

	for k := range story {
		a.Title = story[k].Title
		a.Story = story[k].Story
		a.Options = story[k].Options

		route := fmt.Sprintf("/%s", k)

		http.Handle(route, a)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
