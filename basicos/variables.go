// Declaración e inicialización de variables
// Según el libro "The Go Programming Language", hay cuatro tipos de declaraciones:

// forma corta, sólo para usarse dentro de funciones,
// infiere el tipo por el valor asginado
name := ""

var nombre string // declaración explícita de tipo, usa el valor inicial para el tipo de dato
var nombre = ""   // util para cuando se quieren declarar varias variables

// lo mismo que la anterior, pero es útil para cuando se asigna
// un tipo de dato diferente (cast)
var nombre string = ""