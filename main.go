package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, GitHub Actions")
	log.Printf("Result: %d", Dummy(5))
}

func Dummy(v int) int {
	return v * v
}
