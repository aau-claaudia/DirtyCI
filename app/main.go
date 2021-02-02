package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := atomic.AddUint64(&visitors, 1)

		w.Header().Add("X-Version", Version)

		fmt.Fprintf(w, "Hello %s, this is version %s you are my number %d\n", r.RemoteAddr, Version, v)
	})

	log.Printf("listening on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
