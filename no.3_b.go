package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, 07,30, 00,00,00,00, time.UTC)
	fmt.Println("1)", t)
}