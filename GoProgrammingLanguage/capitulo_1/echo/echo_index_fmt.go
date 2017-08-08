// Echo imprimiendo el índice
package main

import (
	"fmt"
	"os"
)

func main() {
	var sep string

	// en cada iteración declara index y toma el valor de
	// la lista de argumentos
	for index, arg := range os.Args[1:] {
		fmt.Printf(sep+" %d "+arg, index)
		sep = "\n"
	}

}
