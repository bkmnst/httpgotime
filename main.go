package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

var zones = []string{
	"UTC",
	"America/New_York",
	"Europe/London",
	"Asia/Tokyo",
	"Australia/Sydney",
	"Europe/Berlin",
	"America/Los_Angeles",
	"Asia/Kolkata",
	"Africa/Cairo",
	"America/Sao_Paulo",
	"Pacific/Auckland",
	"Europe/Moscow",
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
	rand.Seed(time.Now().UnixNano())
	order := rand.Perm(len(zones))
	for i := 0; i < 10 && i < len(order); i++ {
		name := zones[order[i]]
		loc, err := time.LoadLocation(name)
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "%s: %s\n", name, time.Now().In(loc).Format(time.RFC1123))
	}
	fmt.Fprintf(w, "ARCH: %s\n", runtime.GOARCH)
	fmt.Fprintf(w, "OS: %s\n", runtime.GOOS)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":80", nil)
}
