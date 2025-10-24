package shapes

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type RightAngledTriangle struct {
	Base   float64
	Height float64
}

func (rat RightAngledTriangle) Perimeter() float64 {
	hypotenuse := math.Pow(math.Pow(rat.Base, 2)+math.Pow(rat.Height, 2), 0.5)
	return rat.Base + rat.Height + hypotenuse
}

func (rat RightAngledTriangle) Area() float64 {
	return rat.Base * rat.Height / 2
}
