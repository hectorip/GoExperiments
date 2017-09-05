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
	color.RGBA{255, 10, 200, 1},
	color.RGBA{0, 200, 50, 1},
	color.RGBA{0, 255, 255, 1},
	color.RGBA{0, 200, 50, 1},
}
cl := len(palette) - 1
const (
	whiteIndex = 0 // Colores que se usraán para las imágenes
	blackIndex = 1 // Color de línea
	greenIndex = 2 // Color de línea
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	args := os.Args[1:]
	cycles, _ := strconv.ParseFloat(args[0], 64)
	const (
		res     = 0.001
		size    = 200
		nframes = 64
		delay   = 8
	)
	m, _ := strconv.ParseFloat(args[1], 64)
	freq := rand.Float64() * m
	anim := gif.GIF{LoopCount: nframes} // Creando un GIF
	phase := 0.0
	for i := 0; i < nframes; i++ { // Creando cada cuadro de la animación
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // Se usará como un plano cartesiano
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// fmt.Println(size + int(x*size+0.5))
			index := uint8(rand.Intn(cl) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
