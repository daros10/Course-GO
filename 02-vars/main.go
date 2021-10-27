package main

import "fmt"

// Go es un lenguaje fuertemente tipado, como TS

// Puedes definir variables fuera o dentro de la funcion
// La dieferencia radica en el scope o el alcance de las mismas.
// Mediante := se puede asignar el valor a una varibale sin la palabra reservada var.

// Si a las variables no se les asigna un valor por inicial, GO por defecto al compilar y ejecutar las inicializa, int=0, string="", bool=false
var randomText string
var randomNumber int
var randomBool bool
var textScope string = "text outside scope"

// Notas
// Al Go manejarse con el concepto de paquetes, a veces es necesario porder exportar variables o funciones.
// Para EXPORTAR, basta con poner en MAYUSCULA la primera letra de la varible o de la funcion a exportarse
// Si las variables o funciones estan en MINUSCULA se las considera PRIVADAS

func main() {
	randomNumberWithExplicitType := 4
	randomNumer1, randomNumber2, randomNumber3, randomAwesomeText, textScope := 1, 2, 3, "some awesome text", "text inside scope"

	fmt.Println("Random number with explicit type", randomNumberWithExplicitType)
	fmt.Println("-------------- VALUES --------------")
	fmt.Println(randomNumer1)
	fmt.Println(randomNumber2)
	fmt.Println(randomNumber3)
	fmt.Println(randomNumber)
	fmt.Println(randomAwesomeText)
	fmt.Println(randomText)
	fmt.Println(randomBool)
	fmt.Println("textScopeInside", textScope)
	showTextScope()
	fmt.Println("------------------------------------")
}

func showTextScope() {
	fmt.Println("textScopeOutside", textScope)
}
