package main

import "fmt"

var table [10]int

func main() {
	table[0] = 1
	table[5] = 15
	fmt.Println(table)

	myArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(myArr)

	arr := []int{}
	arr = append(arr, 1)
	arr = append(arr, 2)
	fmt.Println(arr)

	fruits := []string{"apple", "banana", "orange", "coconut", "grap"}
	onlyTwoFruit := fruits[0:2]
	fmt.Println("Only tow fruits", onlyTwoFruit)

	for i := 0; i < len(fruits); i++ {
		if fruits[i] == "banana" {
			fmt.Println("Thre are bananas :P")
		}

		fmt.Println("The rest of fruits are:", fruits[i])
	}

	// tamano 5, capacidad maxima 20
	elements := make([]int, 5, 20)
	fmt.Printf("Size of arr: %d, Capacity of arr: %d\n", len(elements), cap(elements))

	elements = append(elements, 1)
	elements = append(elements, 2)
	elements = append(elements, 3)
	elements = append(elements, 4)
	elements = append(elements, 5)

	for j := 6; j < 21; j++ {
		elements = append(elements, j)
	}

	fmt.Println(elements)

}
