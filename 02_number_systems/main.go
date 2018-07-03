package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d - %b - %x - %q \n", i, i, i, i)
	}
}
