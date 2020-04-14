package main

import (
	"log"
)

// CheckError is a helper function to log errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
