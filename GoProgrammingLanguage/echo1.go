// Imprime los argumentos pasados de regreso en la terminal
package main

// Importar paquete
import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1 ; i <= len(os.Args[1:]); i++ {
        s += sep + os.Args[i]

        // Así se salta el primer loop
        sep = " "
    }
    fmt.Println(s)
}