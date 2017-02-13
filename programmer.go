package main

import "fmt"

type Programmer struct{
    preferredLanguage string
    age int
}

func(brogrammer *Programmer) codes() string {
    return "Codes in " + brogrammer.preferredLanguage
}

func main() {
    var yo Programmer;
    yo.preferredLanguage = "Go";
    fmt.Println(yo.codes())
}