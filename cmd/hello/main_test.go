package main

import (
	"os/exec"
	"testing"
)

func TestVersionFlag(t *testing.T) {
	cmd := exec.Command("go", "run", "./", "--version")
	//cmd.Dir = "cmd/hello"
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("run error:%v", err)
	}
	want := "hello v0.1.0\n"
	if string(out) != want {
		t.Fatalf("run error:%v, want:%v", string(out), want)
	}
}
