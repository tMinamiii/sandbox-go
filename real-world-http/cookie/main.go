package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>2回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>初回訪問</body></html>")
	}

}
