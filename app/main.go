package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	// lovely state
	var visitors uint64

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("unable to determine hostname: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&visitors, 1)
		fmt.Fprintf(w, "Hello %s, this is %s, you are my number %d\n", r.RemoteAddr, hostname, visitors)
	})

	log.Printf("listening on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
