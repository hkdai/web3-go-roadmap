package main

import (
	"io"
	"os"
	"testing"
)

// 1.hello-world
func TestHelloWorld(t *testing.T) {
	want := "hello world\n"
	got := capture(main)
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

//util capture stdout

func capture(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old
	return string(out)
}
