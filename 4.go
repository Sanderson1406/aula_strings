package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = ""
	fmt.Println("Digite uma string")
	fmt.Scanln(&s)
	result := strings.ToUpper(s)
	fmt.Println(result)
}
