// Contador de líneas repetidas sólo de archivos indicados
// Se lee todo el archivo, se cara en memoria y después se procesa
// no se usa "streaming"
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Contando líneas en archivo")
	files := os.Args[1:] // Cada archivo pasado al programa como argumento
	counts := make(map[string]int)
	for _, file := range files {

		data, err := ioutil.ReadFile(file)
		// fmt.Println("Contando líneas del archivo: " + file)
		if err != nil {
			fmt.Printf("Error leyendo archivo %s", file)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") { // seprando por líneas
			counts[line]++
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
