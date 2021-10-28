package main

import "fmt"

// funcion anonima
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

// Las funciones anonimas no tiene nombre
func main() {
	fmt.Println("Sumo 1+2 =", Calculo(1, 2))

	// Resto
	// Eso es lo interesante de las funciones anonimas, al estar alamacenadas en una variable, esta puede ser redefinada
	Calculo = func(i1 int, i2 int) int {
		return i1 - i2
	}
	fmt.Println("Resto 1+2 =", Calculo(2, 1))
	fmt.Println("*************")
	fmt.Println("Calculator Suma:", calculator("+", 5, 5))
	fmt.Println("Calculator Resta:", calculator("-", 5, 4))
	fmt.Println("Calculator Mul:", calculator("*", 5, 5))
	fmt.Println("*************")
	fmt.Println("Closure")
	tableOf := 2
	MyTable := Table(tableOf)
	for i := 1; i < 11; i++ {
		fmt.Println(MyTable())
	}
}

// Funcion que devulver una funcion anonima
func calculator(operation string, num1 int, num2 int) int {
	switch operation {
	case "+":
		resultado := func() int {
			return num1 + num2
		}
		return resultado()
	case "-":
		resultado := func() int {
			return num1 - num2
		}
		return resultado()
	case "*":
		resultado := func() int {
			return num1 * num2
		}
		return resultado()
	default:
		resultado := func() int {
			return num1 / num2
		}
		return resultado()
	}
}

// CLOUSURES
// Permite isolar el codigo (proteger)
func Table(value int) func() int {
	numero := value
	sequence := 0
	return func() int {
		sequence += 1
		return numero * sequence
	}
}
