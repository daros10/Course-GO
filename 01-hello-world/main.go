// En go no existe clases, herencia, polimorfismo, POO
// Go utiliza diferentes formas para solucionar esto
// Es indispensale siempre escribir al inicio de un archivo de go -> package nombre_fichero
package main

// Importaciones, si solo usammos una, se puede poner directo import "fmt"
// Importaciones si usamos muchos paquetes, import ("ftm" "otro" "otro") no es necesario agregar comas
import "fmt"

func main() {
	fmt.Println("Hello world :)")
}

// NOTAS
// Para ejecutar la app basta con: go run main.go
// Si se desea crear un ejecutable: go build main.go -> esto corre atuomaticamente en cualquier SO
