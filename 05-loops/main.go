package main

import "fmt"

// En go no existe el while, ni el do while
// Solo se usa la sentencia FOR

func main() {
	// Different ways to use for

	// 1 way
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("*******")

	// 2 way
	j := 0
	for j < 4 {
		fmt.Println(j)
		j++
	}

	fmt.Println("*******")

	// 3 way - for each
	strings := []string{"hello", "world"}
	for k, s := range strings {
		fmt.Println(k, s)
	}

	for _, words := range strings {
		fmt.Println(words)
	}

	fmt.Println("*******")

	// 4 way - exit loop
	sum := 0
	for i := 1; i < 5; i++ {
		// omit when i takes value equal to 2 without break the loop, so the result OF SUM should be 8
		// continue regresa al ciclo for, goto regresa a la rutina como tal, puede haber mucho mas codigo que puede evaluar antes del loop, esa es la diferencia con goto
		if i == 2 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("*******")

	// 5 way - infinte loop, using break to exit
	secretNumber := 0
	for {
		fmt.Println("Infinite loop")
		fmt.Println("Enter the secret number to break the infinite loop x_x")
		fmt.Scanln(&secretNumber)
		if secretNumber == 10 {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			fmt.Println("Congrats, you can break the infinte loop C:")
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			break
		}

		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("The secret number entered not is in the vault :C, try again ")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	// GOTO
	var m int = 0
SLEEP:
	fmt.Println("Yo goto puedo ejecutar codigo antes del for, cosa que el continue no puede, el solo regresa al loop for")
	for m < 5 {
		if m == 2 {
			m += 2
			fmt.Println("Goto ROUTINE")
			goto SLEEP
		}

		fmt.Printf("Value of m: %d\n", m)
		m++
	}
}
