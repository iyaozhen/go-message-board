package userhandler

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user, I love %s!", r.URL.Path[1:])
}
