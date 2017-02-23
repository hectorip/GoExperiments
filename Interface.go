package main
import "fmt"

type ProgrammerInterface interface{
    codes() string
    your_work() string
}

type Programmer struct{
    preferredLanguage string
}

type Coder struct{
    wear string
    favorite_keyboard string
}

func (p Programmer) codes() string {
    return "Codes in "+ p.preferredLanguage
}

func (p Programmer) your_work() string{
    return "Hello World!"
}

func main() {
    var me Programmer
    me.preferredLanguage = "PHP"
    fmt.Println(me.codes())
}