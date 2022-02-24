package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"hello", "world"}
	for idx, el := range arr {
		fmt.Println(idx, el)
	}

	fmt.Println(strings.Repeat("Hello\n", 10), 34, "hello")
}
