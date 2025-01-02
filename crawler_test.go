package main

import (
    "testing"
	"bytes"
	"os"
)

// TestHelloWorldCapitalization calls crawler.main checking
// for correct capitalization.
func TestHelloWorldCapitalization(t *testing.T) {
	var buf bytes.Buffer
	out := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = out
		w.Close()
	}()
	done := make(chan bool)
	go func() {
		buf.ReadFrom(r)
		r.Close()
		done <- true
	}()

	main()
	w.Close()
	<-done

	// Check the output
	expected := "Hello, World!\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}