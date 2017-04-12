package main

import (
	"fmt"
	"net/http"
	"os"
	"net"
	"strings"
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
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("SplitHostPort returned an error")
		return
	}
	userIp := net.ParseIP(ip)
	if userIp == nil {
		fmt.Println("ParseIP returned nil")
	}

	langList := r.Header.Get("Accept-Language")
	result := strings.Split(langList, ",")
	lang := result[0]

	f := func(c rune) bool {
		return c == '(' || c == ')'
	}
	userAgent := r.Header.Get("User-Agent")
	fields := strings.FieldsFunc(userAgent, f)

	sys := fields[1]

	fmt.Println(userIp)
	fmt.Println(lang)
	fmt.Println(sys)
}
