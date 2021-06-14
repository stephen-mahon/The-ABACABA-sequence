package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(abacaba(5))
}

func abacaba(num int) string {

	alpha := strings.Split("abcdefghijklmnopqrstuvqxyx", "")

	output := ""

	for i := 0; i < num; i++ {
		output = output + alpha[i] + output
	}

	return output
}
