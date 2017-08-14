// Cuenta las líneas repetidas en la entrada de consola
// e imprime la línea con el número de repeticiones
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello repetition counter")
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
