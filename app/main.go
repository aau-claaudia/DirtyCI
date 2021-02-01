package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"sync/atomic"
)

// Version is set by ldflags
var Version string

// ensure Version has at least a non empty value
func init() {
	if Version == "" {
		Version = "not available"
	}
}

func main() {
	// lovely state
	var visitors uint64

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("unable to determine hostname: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitors += 1

		w.Header().Add("X-Version", Version)
		w.Header().Add("X-More-Header", "header")
		fmt.Fprintf(w, "Hello %s, this is %s, you are my number %d\n", r.RemoteAddr, hostname, visitors)
	})

	log.Printf("listening on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
