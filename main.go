package main

import "fmt"

func main() {
	alpha := []byte{97}
	beta := string(alpha[0] + 1)
	fmt.Println(alpha)
	fmt.Println(beta)
}
