// El paquete principal siempre es main, y al existir
// indica que este programa/archivo no es una biblioteca
// sino un programa ejecutable por sí mismo
package main

// No se pueden importar paquetes que no se ocupen
import "fmt"

// La función main indica el punto de entrada de un programa,
// aquí se define lo que el programa hace.
func main() {
    // go es unicode, por lo que puede soportar sin porblema carácteres
    // de todos los kenguajes
	fmt.Println("Hello, 世界!") // Sekai - mundo en japonés 
}
