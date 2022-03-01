package main

import (
	"fmt"
	"log"
	"math"
	"os"

	gd "github.com/misterunix/cgo-gd"
	"github.com/misterunix/colorworks/hsl"
)

func gen3(width, height int, scalefactor, x1start, y1start float64, q1, q2, ff float64,
	min, max float64, huestart, huestop float64, s, l float64) {
	fmt.Println(huestart, huestop)
	// Image width and height
	width = 8192  // width : width of the image
	height = 8192 // height : height of the image

	// Create a new image with the specified width and height.
	//img := image.NewRGBA(image.Rect(0, 0, width, height))
	ibuf0 = gd.CreateTrueColor(width, height)
	bkground := ibuf0.ColorAllocateAlpha(0x00, 0x00, 0x00, 0)
	//white := ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 70)
	//color := ibuf0.ColorAllocateAlpha(255, 208, 0, 100)
	// create a new random number generator with the current time as the seed
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	/*
		zoom := BoundingBox{
			X1: -1.5,
			X2: 1.5,
			Y1: 0.5,
			Y2: 3.5,
		}
	*/

	scalefactor = 2.2
	x1start = -1.0
	y1start = 0.2

	zoom := BoundingBox{
		X1: x1start,
		X2: x1start + scalefactor,
		Y1: y1start,
		Y2: y1start + scalefactor,
	}

	// clear the image with the background color
	ibuf0.FilledRectangle(0, 0, width, height, bkground)

	q1 = 0.690974
	q2 = 0.905823
	ff = 1.746475

	//q1 := rnd.Float64()
	//q2 := rnd.Float64()
	//ff := rnd.Float64() * 3.0

	fmt.Printf("%f %f %f\n", q1, q2, ff)
	fmt.Printf("%+v\n", zoom)

	x := 1.0
	y := 1.0
	//min := 10000000000000.0
	//max := -10000000000000.0
	//var olx1, oly1, d float64
	// 100000000

	for i := 0; i < 100000000; i++ {

		x1 := math.Tan(x)*math.Tan(x) - math.Sin(y)*math.Sin(y) + q1
		y1 := ff*math.Tan(x)*math.Sin(y) + q2

		x2 := convertRange(x1, zoom.X1, zoom.X2, 0, float64(width))
		y2 := convertRange(y1, zoom.Y1, zoom.Y2, 0, float64(height))

		d := math.Sqrt((x1-zoom.X1)*(x1-zoom.X1) + (y1-zoom.Y1)*(y1-zoom.Y1))
		//fmt.Println(d)

		if d > 3 {
			d = 3
		}
		if d < 0 {
			d = 0
		}

		if d < min {
			min = d
		}
		if d > max {
			max = d
		}

		h := convertRange(d, min, max, huestart, huestop)
		s = 1.0
		l = 0.42
		r, g, b := hsl.HSLtoRGB(h, s, l)
		//fmt.Printf("%f %f %d %d %d\n", d, h, r, g, b)
		color := ibuf0.ColorAllocateAlpha(int(r), int(g), int(b), 100)
		ibuf0.SetPixel(int(x2), int(y2), color)
		x = x1
		y = y1

	}
	fmt.Printf("%f %f\n", min, max)
	// Save the image as a PNG.
	imageFilename := fmt.Sprintf("images/%06d.png", os.Getpid()) // set the filename with the pid
	ibuf0.Png(imageFilename)                                     // save the image as a PNG

	// create text file with the image filename
	textFilename := fmt.Sprintf("images/%06d.txt", os.Getpid()) // set the filename with the pid

	infoString1 := fmt.Sprintf("x1:%f y1:%f x2:%f y2:%f scale:%f", zoom.X1, zoom.Y1, zoom.X2, zoom.Y2, scalefactor)
	infoString2 := fmt.Sprintf("q1:%f q2:%f ff:%f", q1, q2, ff)
	infoString3 := fmt.Sprintf("min:%f max:%f", min, max)
	infostring4 := fmt.Sprintf("sat:%f lum:%f", s, l)
	infostring := fmt.Sprintf("%s\n%s\n%s\n%s\n", infoString1, infoString2, infoString3, infostring4)

	f, err := os.Create(textFilename) // create the file
	if err != nil {
		log.Fatal(err) // fail if we can't create the file
	}
	f.WriteString(infostring) // write the image filename to the file
	f.Close()                 // close the file

}
