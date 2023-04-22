package main

import "fmt"

func main() {
	var s string
	fmt.Println("Digite uma string: ")
	fmt.Scanln(&s)
	fmt.Println("A quantide de caracteres nessa string é de", len(s))
}
