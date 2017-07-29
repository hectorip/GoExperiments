// El paquete principal se llama 'main', y al existir
// indica que este programa/archivo no es una biblioteca
// sino un programa ejecutable por sí mismo
package main

// Paquetes importados, para usarse por el programa actual
import "fmt"
// No se pueden importar paquetes que no se ocupen

// La función main indica el punto de entrada de un programa,
// aquí se define lo que el programa hace.
func main() {
    // go es unicode, por lo que puede soportar carácteres
    // de todos los lenguajes
	fmt.Println("Hello, 世界!") 
    // Sekai - mundo en japonés 
}
 // Correr go fmt para formatear