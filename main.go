package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

type Info struct {
	IP       interface{} `json:"ipaddress"`
	Language interface{} `json:"language"`
	Software interface{} `json:"software"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Router)
	http.ListenAndServe(":"+port, nil)
}

func Router(w http.ResponseWriter, r *http.Request) {
	var info = Info{}

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
	software := fields[1]

	info = Info{IP: userIp, Language: lang, Software: software}

	js, err := json.Marshal(info)
	if err != nil {
		fmt.Println("Json Marshal returned nil")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
