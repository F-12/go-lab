package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White,
	color.RGBA{0xbb, 0xcd, 0xc5, 0xff},
	color.RGBA{0xee, 0xde, 0xb0, 0xff},
	color.RGBA{0xea, 0xff, 0x56, 0xff},
	color.RGBA{0x40, 0xde, 0x5a, 0xff},
	color.RGBA{0xbc, 0xe6, 0x72, 0xff},
	color.RGBA{0xff, 0x4c, 0x00, 0xff},
	color.RGBA{0xff, 0x46, 0x1f, 0xff},
	color.RGBA{0x42, 0x50, 0x66, 0xff},
	color.RGBA{0x70, 0xf3, 0xff, 0xff}}

const (
	whiteIndex = 0 // first color in palette
)

var r = rand.New(rand.NewSource(42))

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0 // phase difference
	idxp := len(palette)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%idxp))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
