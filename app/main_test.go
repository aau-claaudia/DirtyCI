package main

import (
	"net/http"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {

	go main()

	var wg sync.WaitGroup

	// hammer the service with 100 requests, expect no race conditions
	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			resp, err := http.Get("http://localhost:80")
			if err != nil {
				t.Fatalf("cannot GET url: %s", err)
			}

			resp.Body.Close()

			wg.Done()
		}()
	}
	wg.Wait()
}
