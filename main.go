package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", args[1])
	s, err := getHTML(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s)
}
