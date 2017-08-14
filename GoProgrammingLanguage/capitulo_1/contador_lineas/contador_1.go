// Cuenta las líneas repetidas en la entrada de consola
// e imprime la línea con el número de repeticiones
package main

import (
	"fmt"
	"bufio"
)

func main() {
	fmt.Println("Hello repetition counter")
	var scanner = bufio.NewScanner();
	for line := scanner.Input() {
		fmt.Println(line)
	}
}
