package main

import (
	"fmt"
	"strconv"

	"github.com/vinicio-ochoa/gocourse/variables/funciones"
)

func main() {
	fmt.Printf("hello, world\n")
	var resultado = funciones.Sub(10, 20)
	var suma = strconv.Itoa(resultado)
	fmt.Println("La suma es " + suma)
	fmt.Println("La suma es " + funciones.VariableCapitalizada)
}
