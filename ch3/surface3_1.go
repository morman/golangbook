// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axes ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixles per x or y unit
	zscale        = height * 0.4        // pixles per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// if any of these coordinates are not a number, skip
			validCoord := true

			// build an array of the coordinates
			coordarray := [...]float64{ax, ay, bx, by, cx, cy, dx, dy}

			// Loop over coordinates
			// Note: range returns both the index and value: we don't need the index
			for _, v := range coordarray {
				// Is it positive infinity?
				if math.IsInf(v, 1) {
					validCoord = false
				}
				// Is it negative infinity?
				if math.IsInf(v, -1) {
					validCoord = false
				}
				// Is it either infinity?
				if math.IsInf(v, 0) {
					validCoord = false
				}
				// Is it not a number?
				if math.IsNaN(v) {
					validCoord = false
				}

			}

			// Test if all the coords are valid float64
			if validCoord == true {
				// create a polygon
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				// render a comment
				fmt.Printf("<!-- INVALID POINTS: '%g,%g,%g,%g,%g,%g,%g,%g' -->\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// go build surface3_1.go;./surface3_1 > surface3_1.svg; open -a safari surface3_1.svg
