package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var c string
	fmt.Print("Digite uma String: ")
	fmt.Scanln(&s)
	fmt.Print("Digite um Caractére: ")
	fmt.Scanln(&c)
	result := strings.Count(s, c)
	fmt.Println(result)
}
