package main

import (
	"fmt"
	"strings"
)

func ReverseString(sentenses string) string {
	tmp := strings.Split(sentenses, "")
	data := ""

	for i := len(sentenses) - 1; i >= 0; i-- {
		data += tmp[i]
	}
	return data
}

func main()  {
	s := ReverseString("Aplikasi")
	fmt.Println(s)
}