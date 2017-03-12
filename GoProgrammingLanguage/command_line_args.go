package main

import (
	"fmt"
	"os"
)

// la función main es obligatoria
func main() {
	fmt.Println("Este es un programa que lee los argumentos de consola: \n")
	fmt.Println(os.Args)

	// Args es un slice, que en Go se comportan como una lista
	cml_arguments := os.Args[1:]
	fmt.Println(cml_arguments)
}
