// Lissajous generates GIF animations
// of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xff, 0xff, 0x10, 0xff},
	color.RGBA{0xff, 0x10, 0x10, 0x11},
	color.RGBA{0x10, 0xff, 0x10, 0xff},
	color.RGBA{0x10, 0x10, 0xff, 0x11}}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	yellowIndex = 2
	redIndex    = 3
	greenIndex  = 4
	blueIndex   = 5
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 16     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.2 // phase difference
	var currentColor uint8 = 0
	var currentColorFrames uint8 = 0
	var framesStep uint8 = 96
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// switch color after some frames
			currentColorFrames++
			if currentColorFrames > framesStep {
				currentColorFrames = 0
				currentColor++
				if currentColor > blueIndex {
					currentColor = blackIndex
				}
			}

			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				currentColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: Ignoring encoding errors
}
