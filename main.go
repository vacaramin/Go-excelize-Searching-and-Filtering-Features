package main

import (
	"fmt"
	"net/http"
	"os"
	handlers "search/src/Handlers"
	"search/src/initilaizers"
)

func main() {
	initilaizers.Init()

	// routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/process-excel", handlers.ProcessExcel)

	fmt.Println("listening and serving on Port", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
