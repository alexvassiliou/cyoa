package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

//Story is a collection of adventures
type Story map[string]Adventure

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

// Parse the json input into an adventure
func Parse(filePath string) Story {
	in, err1 := ioutil.ReadFile(filePath)
	CheckError(err1)

	var a Story

	err2 := json.Unmarshal(in, &a)
	CheckError(err2)

	return a
}
