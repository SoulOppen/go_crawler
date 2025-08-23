package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")
	for k, v := range pages {
		fmt.Printf("Found %d internal links to %s\n", v, k)
	}
}
