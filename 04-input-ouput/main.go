package main

import (
	"bufio"
	"fmt"
	"os"
)

var num1 int
var num2 int
var desc string

func main() {
	fmt.Println("Enter number 1: ")
	fmt.Scanf("%d", &num1)

	fmt.Println("Enter number 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Println("Number entered are:", num1, num2)

	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter description: ")

	if scan.Scan() {
		desc = scan.Text()
	}

	fmt.Println("Description entered is:", desc)

}
