// Este programa puede contar las líneas proveninentes de un
// archivo o de la línea de comandos
package contador_lineas

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // Todos los argumentos que se pasaron al programa
	if len(files) == 0 {
		// No se pasaron nombres de archivo al programa
		// Se pasa a countLines la entrada estándar (le línea de comandos)
		// como la fuente del text a examinar
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				continue // Termina esta iteración del ciclo
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	// Cuenta las líneas efectivamente y las guarda en el
	// diccionario que se le pasó

	// Todos los diccionarios son pasados por "referencia"
	// lo que indica que los cambios hechos en esta función
	// afectarán la variable externa

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
