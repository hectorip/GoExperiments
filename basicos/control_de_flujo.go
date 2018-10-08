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

	s := rand.NewSource(time.Now().UnixNano()) // poniendo una semilla más o menos aleatoria
	r := rand.New(s)

	// Se acepta una declaración o llamada de función, pero sólo vive dentor del ámbito del
	// if o else
	if count := r.Intn(100); count > 10 {
		fmt.Printf("El número es mayor que 10: %d\n", count)
	} else {
		fmt.Printf("El número no es mayor que 10: %d\n", count)
	}

	// Switch
	switch count := r.Intn(1000); count % 2 { // también acepta una declaracón antes
	case 0:
		fmt.Printf("El número es par %d\n", count)
	case 1:
		fmt.Printf("El número es impar %d\n", count)
	}

	// Tagless Switch
	// No necesita una expresión para comparar y ejecuta el primer bloque que cumpla
	// con las condiciones
	x := r.Intn(20)
	switch {
	case x < 5:
		fmt.Printf("El número es menor que 5: %d", x)
	case x <= 10:
		fmt.Printf("El número está entre 5 y 10: %d", x)
	case x > 10:
		fmt.Printf("El número es mayor que 10: %d", x)
	}

}

//If
