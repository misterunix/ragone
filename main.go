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

	//gen2()

	/*
		x1:-1.000000 y1:0.200000 x2:1.200000 y2:2.400000 scale:2.200000
		q1:0.690974 q2:0.905823 ff:1.746475
		min:0.819748 max:3.000000
		sat:1.000000 lum:0.420000

	*/

	gen3(8192, 8192,
		2.2, -1.0, 0.2,
		0.690974, 0.905823, 1.746475,
		0.819748, 3.000000,
		234, 10,
		1.0, 0.420000)

}

// convert a number range to another number range
func convertRange(value float64, oldMin float64, oldMax float64, newMin float64, newMax float64) float64 {
	return (((value - oldMin) * (newMax - newMin)) / (oldMax - oldMin)) + newMin
}
