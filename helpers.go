package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// CheckError is a helper function to log errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Parse the json input into an adventure
func Parse(filePath string) map[string]Adventure {
	in, err1 := ioutil.ReadFile(filePath)
	CheckError(err1)

	var a map[string]Adventure

	err2 := json.Unmarshal(in, &a)
	CheckError(err2)

	return a
}
