package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Router)
	http.ListenAndServe(":"+port, nil)
}

func Router(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	lang := r.Header.Get("Accept-Language")
	sys := r.Header.Get("User-Agent")
	fmt.Println(ip)
	fmt.Println(lang)
	fmt.Println(sys)
}
