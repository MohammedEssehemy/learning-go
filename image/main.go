// package main

// import (
// 	"fmt"
// 	"image"
// )

// func main() {
// 	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	fmt.Println(m.Bounds())
// 	fmt.Println(m.At(0, 0).RGBA())
// }

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x int
	y int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.x, img.y)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 65}
	pic.ShowImage(m)
}
