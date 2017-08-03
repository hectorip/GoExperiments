// Imprime los argumentos pasados de regreso en la terminal
package main

// Los imports pueden ser una colección
// entre paréntesis
import (
	"fmt"
	"os"
)

func main() {
	// Declaración de variables con tipo explícito
	// al no asignarles valores inicialmente, se
	// establecen como cadenas vacías("") para
	// los strings.
	var s, sep string

	// Ciclo `for`
	// Consiste de trees partes, iniciialización,
	// prueba y post-operación
	for i := 1; i <= len(os.Args[1:]); i++ { // no se usan paréntesis para la configuración del ciclog
		s += sep + os.Args[i]

		// Saltando el primer seprador
		sep = " "
	}
	fmt.Println(s)
}
