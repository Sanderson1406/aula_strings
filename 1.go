package main

import "fmt"

func main() {
	var s1, s2 string
	var s3 string
	fmt.Println("Digite uma string: ")
	fmt.Scanln(&s1)
	fmt.Println("Digite outra string: ")
	fmt.Scanln(&s2)
	s3 = s1 + ", " + s2
	fmt.Println(s3)
}
