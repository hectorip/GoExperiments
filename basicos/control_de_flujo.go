// Control de flujo en Go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	if true { // expresión booleane
		fmt.Println("Esta condición siempre se cumple")
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	if count := r.Intn(100); count > 10 { // Se acepta una declaración o llamada de función
		fmt.Printf("El número es mayor que 10: %d", count)
	} else {
		fmt.Printf("El número no es mayor que 10: %d", count)
	}
}

//If
