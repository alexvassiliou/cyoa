package main

import (
	"fmt"
	"net/http"
)

var storyHandler Adventure

func main() {
	filePath := "gopher.json"

	adventures := Parse(filePath)

	for k := range adventures {
		storyHandler.Title = adventures[k].Title
		storyHandler.Story = adventures[k].Story
		storyHandler.Options = adventures[k].Options

		route := fmt.Sprintf("/%s", k)

		http.Handle(route, storyHandler)
	}

	err := http.ListenAndServe(":8080", nil)
	CheckError(err)

}
