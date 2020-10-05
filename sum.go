package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("5 + 5 = %d", sum(5, 5))
}