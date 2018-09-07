package main

import (
	"fmt"
	"strconv"

	"github.com/vinicio-ochoa/gocourse/variables/funciones"
)

func main() {
	fmt.Printf("hello, world\n")
	var resultado = funciones.Sub(20, 2)
	var suma = strconv.Itoa(resultado)
	fmt.Println(suma)
}
