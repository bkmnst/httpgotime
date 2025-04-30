package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n", time.Now())
	fmt.Fprintf(w, "ARCH: %s\n", runtime.GOARCH)
	fmt.Fprintf(w, "OS: %s\n", runtime.GOOS)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8081", nil)
}
