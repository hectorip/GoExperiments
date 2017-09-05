// Crea un Gif con una figura de Lissajous
// (curvas extrañas)
package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	// "sys"
	"strconv"
)

// Declaramos la paleta de Colores que usaremos
// El primero lo usaremos como fondo y el segundo como color de líneas
var palette = []color.Color{
	color.White,
	color.RGBA{255, 10, 100, 1},
	color.RGBA{0, 200, 50, 1},
	color.RGBA{255, 0, 0, 1},
	color.RGBA{0, 255, 0, 1},
	color.RGBA{0, 0, 255, 1},
}
var cl = len(palette) - 1

const (
	whiteIndex = 0 // Colores que se usraán para las imágenes
	blackIndex = 1 // Color de línea
	greenIndex = 2 // Color de línea
)

func main() {
	// fmt.Println(cl)
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	args := os.Args[1:]
	cycles, _ := strconv.ParseFloat(args[0], 64)
	const (
		res     = 0.0001
		size    = 200
		nframes = 128
		delay   = 8
	)
	m, _ := strconv.ParseFloat(args[1], 64)
	freq := rand.Float64() * m
	anim := gif.GIF{LoopCount: nframes} // Creando un GIF
	phase := 0.0
	for i := 0; i < nframes; i++ { // Creando cada cuadro de la animación
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // Se usará como un plano cartesiano
		img := image.NewPaletted(rect, palette)
		var index = uint8(1)
		var t2 float64
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// fmt.Println(size + int(x*size+0.5))
			t2 += res
			// if 3.1415-t2 <= 0.1 {
			// 	// fmt.Println("YESS")
			// 	index = uint8(rand.Intn(cl) + 1)
			// 	// fmt.Println(index)
			// 	t2 = 0
			// }
			if y >= 0.0 && y <= 0.3 {
				index = 5
			} else if y >= 0.3 && y <= 0.6 {
				index = 1
			} else if y >= 0.6 && y <= 0.9 {
				index = 2
			} else {
				index = 0
			}
			// fmt.Println(index)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.2
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
