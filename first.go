package main
import "fmt"

type Persona struct{
	Nombre string
}
func main() {
    var yo Persona;
    yo.Nombre = "Héctor";
    fmt.Println("Hello " + yo.Nombre)
    fmt.Print("nais")
    fmt.Println("perrón")
}

 