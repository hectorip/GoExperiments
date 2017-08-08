package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Este es un programa que lee los argumentos de consola: \n")
	fmt.Println(os.Args)

	// Args es un slice, que en Go se comportan como una lista
	cml_arguments := os.Args[1:] // Se elimina el primer argumento que es el nombre del programa mismo
	fmt.Println(cml_arguments)
	// Por el momento, no me gusta nada que la identación sean tabs.
}
