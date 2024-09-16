package main

import (
	"fmt"

	"projeto-go/meupacote"
)

func main() {
	saudacao := meupacote.Saudar("Mundo")
	fmt.Println(saudacao)
}
