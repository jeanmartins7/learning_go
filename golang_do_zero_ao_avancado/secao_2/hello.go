package hello_world

import "fmt"

func init() {
	salario := 999.00
	var salarioMaisBonus float64

	salarioMaisBonus = salario

	if salario < 1000 {
		salarioMaisBonus = salario + 100
	}

	fmt.Println("Salário: ", salarioMaisBonus)

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	ii := 0

	for true {
		if ii > 10 {
			break

		}

		fmt.Println(ii)

		ii += 1
	}

	var moradia []string
	moradia = append(moradia, "casa")
	moradia = append(moradia, "apartamento")
	moradia = append(moradia, "rancho")

	for ii := range moradia {
		fmt.Println(moradia[ii])
	}

	for ii, jj := range moradia {
		fmt.Print(ii)
		fmt.Println(jj)
	}

	ultimo := moradia[len(moradia)-1:]
	fmt.Print(ultimo)

	m := make(map[int]int)
	m[1] = 10
	fmt.Print(m[1])

	hello()

}

func hello() {
	fmt.Println("hello world")
}

func soma(numero1 int, numero2 int) int {
	return numero1 + numero2
}
