package main

import (
	"html/template"
	"net/http"
)

// Adventure maps each story element of the json
type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func (a Adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("story.html")
	CheckError(err)

	err2 := t.Execute(w, a)
	CheckError(err2)
}
