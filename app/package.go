package main

import "fmt"

func init() {
	fmt.Printf("%s\n", "package inited")
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
