package main

import "fmt"

func main() {
	fmt.Println(one(10))
	fmt.Println("****************")

	// En go se debe almacenar en variables el valor que retonrna la funcion
	number, flag := two(10)

	fmt.Println("Number", number)
	fmt.Println("flag", flag)
	fmt.Println("****************")

	fmt.Println(Calculo(1, 2))
	fmt.Println("****************")
	fmt.Println(Calculo(1, 2, 3, 4, 5))
	fmt.Println("****************")
	fmt.Println(Calculo(1, 2, 45, 6456, 7467, 564))
}

func one(number int) int {
	return number * 2
}

// En GO una funcion puede devolver mas de un tipo
func two(number int) (int, bool) {
	if number == 2 {
		return number, false
	}

	return number, true
}

// En go se pueden crear una funcion que reciba un parametro sin especificar cuantos sean
// es decir el parametro number puede ser, 1,2,3 o n numeros, eso se lo consigue con los 3 puntos ...
func Calculo(number ...int) int {
	total := 0

	for i, num := range number {
		total += num
		fmt.Printf("\n Item %d\n", i)
	}

	return total

}
