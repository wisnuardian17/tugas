package main

import "fmt"

func Hello() {
	fmt.Println("Hello")
}

func main() {
	for i := 0; i < 3; i++ {
		Hello()
	}
}
