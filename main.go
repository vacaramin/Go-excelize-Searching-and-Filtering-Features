package main

import (
	"fmt"
	"net/http"
	"os"
	controllers "xlsx/src/Controllers"
	"xlsx/src/initilaizers"
)

func main() {
	initilaizers.Init()

	// routes
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/process-excel", controllers.ProcessExcel)

	fmt.Println("listening and serving on Port", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
