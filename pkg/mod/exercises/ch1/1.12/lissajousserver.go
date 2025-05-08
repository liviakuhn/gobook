// Lissajous server generates GIF animations of random Lissajous
// figures.
package main

import (
	"strconv"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"log"
	"net/http"
	"time"
)

var palette = []color.Color{
		color.White,
		color.Black,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0  // first color in palette
	blackIndex = 1 // next color in palette
	redIndex = 2
	greenIndex = 3
	blueIndex = 4
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
}

// handler displays a Lissajous figure.
func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	params := make(map[string]float64)
	for k, v := range r.URL.Query() {
		numValue, err := strconv.ParseFloat(v[0], 64)
		// numValue, err := strconv.ParseFloat(v.get(k), 64)
		params[k] = numValue
		if (err != nil) {
			log.Print(err)
		}
	}
	lissajous(w, params)
}

func lissajous(out io.Writer, requestParams map[string]float64) {
	defaultParams := map[string]float64{
    		"cycles": 5,
		"res": 0.001,
		"size": 100,
		"nframes": 64,
		"delay": 8,
	}
	params := make(map[string]float64)
	for k, v := range defaultParams {
		value, ok := requestParams[k]
		if (!ok || value <= 0) {
			params[k] = v
		} else {
			params[k] = value
		}
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(params["nframes"])}
	phase := 0.0 // phase difference
	for i := 0; i < int(params["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*int(params["size"])+1, 2*int(params["size"])+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < params["cycles"]*2*math.Pi; t += params["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(rand.Intn(blueIndex - blackIndex + 1) + blackIndex) // n ∈ [blackIndex, blueIndex]
			img.SetColorIndex(int(params["size"])+int(x*params["size"]+0.5), int(params["size"])+int(y*params["size"]+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(params["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
