package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		fmt.Fprintf(w, "Entorno: %s | Servidor: %s\n", os.Getenv("ENTORNO"), host)
	})
	http.ListenAndServe(":8080", nil)
}