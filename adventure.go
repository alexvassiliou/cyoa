package main

import (
	"io"
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
	io.WriteString(w, "hello from baz")
}
