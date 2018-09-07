package main

import (
	"flag"
	"image"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

var (
	input     = flag.String("i", "", "Image to input into algorithm")
	out       = flag.String("o", "out", "Name of image to output without extension")
	passes    = flag.Uint("p", 5, "Number of passes")
	attempts  = flag.Uint("a", 10, "Number of attempts to try swapping")
	colorFrom = flag.String("cs", "", "Color source, must be same size as input, defaults to input")
	smooth    = flag.Uint("smooth", 0, "Number of times to smooth out the image")
	silent    = flag.Bool("silent", false, "Silences any logging")
	cprofile  = flag.String("cpuprofile", "", "File to write cpuprofile to")
)

var outExt string

func must(o interface{}, err error) interface{} {
	notNilError(err)
	return o
}

func notNilError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func LoadImage(src string) image.Image {
	f := must(os.Open(src)).(*os.File)
	defer f.Close()
	img, _, err := image.Decode(f)
	notNilError(err)
	return img
}

func SaveImage(to string, src image.Image) {
	output := must(os.Create(to)).(*os.File)
	defer output.Close()

	err := png.Encode(output, src)
	if err != nil {
		log.Fatalln(err)
	}
}

func Setup() {
	log.SetFlags(log.Ltime)
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	switch {
	case *input == "":
		log.Fatalln("Must pass input flag \"-i\" to image")
	case *colorFrom == "":
		*colorFrom = *input
	}
	outExt = filepath.Ext(*input)
	if *silent {
		log.SetOutput(ioutil.Discard)
	}
	if filepath.Ext(*input) == ".gif" {
		HandleGif()
		os.Exit(0)
	}
}

func main() {
	Setup()
	if *cprofile != "" {
		cpufile := must(os.Open(*cprofile)).(*os.File)
		pprof.StartCPUProfile(cpufile)
		defer pprof.StopCPUProfile()
	}
	img := LoadImage(*input)
	colorImg := LoadImage(*colorFrom)
	SaveImage(*out+".png", Sort(img, colorImg))
}

func Sort(design, colors image.Image) image.Image {
	bounds := design.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	colors = resize.Resize(uint(w), uint(h), colors, resize.Bicubic)

	cb := NewColorBlock(colors)

	for p := uint(0); p < *passes; p++ {
		log.Printf("Starting pass %d \n", p)
		for x := w - 1; x > 0; x-- {
			for y := h - 1; y > 0; y-- {
				for a := uint(0); a < *attempts; a++ {
					if cb.TrySwap(x, y, rand.Intn(x), rand.Intn(y), design) {
						break
					}
				}
			}
		}
	}

	imgOut := cb.ToImage()
	for i := uint(0); i < *smooth; i++ {
		imgOut = Smooth(imgOut)
	}
	return imgOut
}

func HandleGif() {
	f := must(os.Open(*input)).(*os.File)
	defer f.Close()
	g := must(gif.DecodeAll(f)).(*gif.GIF)
	w, h := g.Config.Width, g.Config.Height

	colorImg := LoadImage(*colorFrom)
	colorImg = resize.Resize(uint(w), uint(h), colorImg, resize.Bicubic)
	palette := GetPalette(colorImg)

	log.Printf("Running on gif with %d images\n", len(g.Image))
	work := make(chan struct {
		int
		*image.Paletted
	}, len(g.Image))

	for i, v := range g.Image {
		work <- struct {
			int
			*image.Paletted
		}{i, v}
	}
	close(work)

	var wg sync.WaitGroup

	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for w := range work {
				log.Printf("Handling image %d\n", w.int)
				g.Image[w.int] = Palettize(Sort(w.Paletted, colorImg), palette)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	output := must(os.Create(*out + ".gif")).(*os.File)
	defer output.Close()
	notNilError(gif.EncodeAll(output, g))
}
