package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		r.URL.Query().Get("get")
		r.Form.Get("form")
		io.WriteString(w, "Hello world!")
	})
	http.ListenAndServe(":8080", mux)
}
