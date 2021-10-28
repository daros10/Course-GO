package main

import "fmt"

var status string = "pennding"

func main() {

	if status == "pennding" {
		fmt.Println("Value of status:", status)
	} else {
		fmt.Println("Value of status diffetent to pennding:", status)
	}

	// Evaluar y asignar valor a la variable en la misma linea del if
	var status2 string = "completed"

	if status2 = "CHANGED TO A NEW VALUE IF SENTECES EVALUATED IS TRUE IN SINGLE LINE"; status2 == "completed" {
		fmt.Println("-Value of status changet to:", status2)
	} else {
		fmt.Println("-Value of status diffetent to completed:", status2)
	}

	// switch
	switch number := 5; number {
	case 1:
		fmt.Println("Num", number)
	case 2:
		fmt.Println("Num", number)
	case 3:
		fmt.Println("Num", number)
	case 4:
		fmt.Println("Num", number)
	case 5:
		fmt.Println("Num", number)
	default:
		fmt.Println("Num major to five")
	}
}
