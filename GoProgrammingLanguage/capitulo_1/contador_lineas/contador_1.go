// Cuenta las líneas repetidas en la entrada de consola
// e imprime la línea con el número de repeticiones
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Contaré las líneas repetidas que me mandes. Escríbelas y cuando termines presiona `CTRL + D`\n")
	repetitions := make(map[string]int)
	// construyendo un scanner de la entrada estándar (consola)
	var scanner = bufio.NewScanner(os.Stdin)

	// Este for funciona exactamente como un while
	for scanner.Scan() {
		// Si la entrada no existe, inmediatamente se crea con 0,
		// por ser el valor inicial default de los enteros, por lo
		// que podemos incrementarlo con `++` sin ningún problema.
		repetitions[scanner.Text()]++
	} // El scanner devuelve falso cuando encuentra EOF,
	// en la consola se manda con CTRL + D

	// Mostrando las líneas que se repiten, recorriendo el arreglo y
	// filtrando las que no tienen en su cuenta más que dos
	fmt.Println("\n\nLíneas repetidas: \n")
	for line, reps := range repetitions {
		if reps > 1 {
			fmt.Printf("%s \t %d \n", line, reps)
		}
	}
}
