package source

import (
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Helloworld")
}
