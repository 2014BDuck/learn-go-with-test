// @Author: 2014BDuck
// @Date: 2021/2/14

package shapes

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Length)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Length
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (tri Triangle) Area() float64 {
	return tri.Base * tri.Height / 2
}
