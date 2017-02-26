// Tercera versión usando el paquete strings
package main

import (
    "fmt"
    "os"
    "strings"
)

func main(){
    fmt.Println(strings.Join(os.Args[1:], " "))
}