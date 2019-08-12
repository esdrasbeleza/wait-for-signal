package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Helo world!")
	})

	name, _ := os.Hostname()

	fmt.Printf("Server running at http://%s:80\n", name)
	http.ListenAndServe(":80", nil)
}
