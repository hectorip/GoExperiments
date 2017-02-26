// Echo imprimiendo el índice
package main

import (
    "strconv"
    "fmt"
    "os"
)

func main(){
    var s, sep string

    for  index, arg := range os.Args[1:] {
        s += sep + strconv.Itoa(index) + " " + arg
        sep = "\n"
    }
    fmt.Println(s)
}