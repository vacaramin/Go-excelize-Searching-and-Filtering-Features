package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Helloworld")
}
