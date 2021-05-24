package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "index, I love %s!", r.URL.Path)
	}
}
