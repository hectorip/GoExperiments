// Imprime los argumentos pasados de regreso en la terminal
package main

// Los imports pueden ser una colección
// entre paréntesis
import (
	"fmt"
	"os"
)

func main() {
    // declaración de variables con tipo explícito
    // al no asignarles valores inicialmente, se van
    // a "", para los strings.
	var s, sep string
	for i := 1; i <= len(os.Args[1:]); i++ {
		s += sep + os.Args[i]

		// Así se salta el primer loop
		sep = " "
	}
	fmt.Println(s)
}
