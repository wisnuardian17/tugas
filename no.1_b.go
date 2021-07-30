package main

import "fmt"

func World() {
	fmt.Println("World")
}

func main() {
	for i := 0; i < 5; i++ {
		World()
	}
}
