package main
import "fmt"

type ProgrammerInterface interface{
    codes() string
}

type Programmer struct{
    preferredLanguage string
}

func (p Programmer) codes() string {
    return "This guy codes in "+ p.preferredLanguage
}

func main() {
    var me Programmer
    me.preferredLanguage = "PHP"
    fmt.Println(me.codes())
}