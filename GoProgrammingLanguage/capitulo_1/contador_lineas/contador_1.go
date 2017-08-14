// Cuenta las líneas repetidas en la entrada de consola
// e imprime la línea con el número de repeticiones
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello repetition counter \n")
	repetitions := make(map[string]int)
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		repetitions[scanner.Text()]++
	}

	for line, reps := range repetitions {
		if reps > 1 {
			fmt.Printf("%s \t %d\n", line, reps)
		}
	}
}
