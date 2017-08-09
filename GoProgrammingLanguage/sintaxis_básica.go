// Sintaxis Básica
// Este archivo ejemplefica la sintaxis básica de Go, 
// sin embargo, no es ejecutable, no contiene la eestructura
// básica de un archivo de Go y repite nombres de
// variable para hacer más entendible.

// Creación de variables

var name string // name = "" | "" es el valor default de los strings 
name := "" // name es un tipo string, al habérsele asignado una cadena
var name = "" // Se infiere el tipo, a string, por la asignación
var name string = "" // redundante, no debería usarse muy seguido

// Tipos de dato

var age int
var gender bool
var name string
var speed float64

// Estructuras (structs)

// Son tipos de dato compuestos diseñados para representar información
// compleja, pueden pensarse como la parte de las propiedades
// de los objetos en los lenguajes que lo soportan.

// 
struct {
  Name string
  Age int // los campos con nombres en mayúscula son públicos
  gender bool  
}



