package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my websiteasd! sdfsdf sdf")
  })
  http.HandleFunc("/test", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World xx")
	})
	http.ListenAndServe(":80", nil)
}
