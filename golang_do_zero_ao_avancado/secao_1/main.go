package main

import "fmt"

func main() {
	var peso int
	peso = 10
	var nome string
	nome = "Jean"
	fmt.Println("hello, teste")
	fmt.Println(peso)
	fmt.Println(nome)

	mensagem := "test"
	fmt.Println(mensagem)

	// lista := ["jean", "ana", "joao"]
	// fmt.Println(lista[0])

	//tipo primitivo
	texto := "string" //string
	numero := 10      //int
	metro := 1.5      //float32
	ehFalso := false  //bool

	fmt.Println(texto)
	fmt.Println(numero)
	fmt.Println(metro)
	fmt.Println(ehFalso)

	constante()
}
