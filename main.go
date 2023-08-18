package main

import (
	"net/http"
	"os"
	"search/src/initilaizers"
)

func main() {
	initilaizers.Init()
	// http.HandleFunc("/", fm)
	//
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
