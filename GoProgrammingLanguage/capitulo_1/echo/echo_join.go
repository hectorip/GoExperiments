// La manera más corta de implementar
// el echo, usando funciones predefinidas de Go

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
