/*
BSD 3-Clause License

Copyright (c) 2022, William Jones
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	gd "github.com/misterunix/cgo-gd"
)

// struct to hold bounding box data
type BoundingBox struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

var ibuf0 *gd.Image

func main() {

	// Image width and height
	width := 1024  // width : width of the image
	height := 1024 // height : height of the image

	// Create a new image with the specified width and height.
	//img := image.NewRGBA(image.Rect(0, 0, width, height))
	ibuf0 = gd.CreateTrueColor(width, height)
	bkground := ibuf0.ColorAllocateAlpha(0x00, 0x00, 0x00, 0)
	white := ibuf0.ColorAllocateAlpha(0xFF, 0xFF, 0xFF, 70)

	// create a new random number generator with the current time as the seed
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	zoom := BoundingBox{
		X1: -5.0,
		Y1: -5.0,
		X2: 5.0,
		Y2: 5.0,
	}

	for count := 0; count < 100; count++ {

		// clear the image with the background color
		ibuf0.FilledRectangle(0, 0, width, height, bkground)

		q1 := rnd.Float64()
		q2 := rnd.Float64()
		ff := rnd.Float64() * 3.0

		fmt.Printf("%f %f %f\n", q1, q2, ff)
		fmt.Printf("%+v\n", zoom)

		x := 1.0
		y := 1.0
		for i := 0; i < 100000; i++ {

			x1 := math.Tan(x)*math.Tan(x) - math.Sin(y)*math.Sin(y) + q1
			y1 := ff*math.Tan(x)*math.Sin(y) + q2

			x2 := convertRange(x1, zoom.X1, zoom.X2, 0, float64(width))
			y2 := convertRange(y1, zoom.Y1, zoom.Y2, 0, float64(height))

			ibuf0.SetPixel(int(x2), int(y2), white)

			x = x1
			y = y1
		}

		// Save the image as a PNG.
		imageFilename := fmt.Sprintf("images/%06d-%04d.png", os.Getpid(), count) // set the filename with the pid
		ibuf0.Png(imageFilename)                                                 // save the image as a PNG

		// create text file with the image filename
		textFilename := fmt.Sprintf("images/%06d-%04d.txt", os.Getpid(), count) // set the filename with the pid

		infoString1 := fmt.Sprintf("x1:%f y1:%f x2:%f y2:%f", zoom.X1, zoom.Y1, zoom.X2, zoom.Y2)
		infoString2 := fmt.Sprintf("q1:%f q2:%f ff:%f", q1, q2, ff)
		infostring := fmt.Sprintf("%s\n%s\n", infoString1, infoString2)

		f, err := os.Create(textFilename) // create the file
		if err != nil {
			log.Fatal(err) // fail if we can't create the file
		}
		f.WriteString(infostring) // write the image filename to the file
		f.Close()                 // close the file

	}

}

// convert a number range to another number range
func convertRange(value float64, oldMin float64, oldMax float64, newMin float64, newMax float64) float64 {
	return (((value - oldMin) * (newMax - newMin)) / (oldMax - oldMin)) + newMin
}
