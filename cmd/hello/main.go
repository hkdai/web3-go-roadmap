package main

import (
	"flag"
	"fmt"
)

var version = "v0.1.0"

func main() {
	showVer := flag.Bool("version", false, "print version")
	flag.Parse()
	if *showVer {
		fmt.Println("hello", version)
	}
	//fmt.Println("hello,web3!")
}
