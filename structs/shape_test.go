package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(tb testing.TB, shape Shape, expected float64) {
		tb.Helper()
		got := shape.Perimeter()
		if got != expected {
			t.Errorf("got %.2f but expected %.2f from %#v", got, expected, shape)
		}
	}

	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{
			shape: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			expected: 40.0,
		},
		{
			shape: Circle{
				Radius: 7.0,
			},
			expected: 2 * math.Pi * 7.0,
		},
		{
			shape: RightAngledTriangle{
				Base:   3.0,
				Height: 4.0,
			},
			expected: 12.0,
		},
	}

	for _, pt := range perimeterTests {
		checkPerimeter(t, pt.shape, pt.expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(tb testing.TB, shape Shape, expected float64) {
		tb.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got %.2f but expected %.2f from %#v", got, expected, shape)
		}
	}

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{10.0, 10.0}, expected: 100.0},
		{shape: Circle{7.0}, expected: math.Pi * math.Pow(7.0, 2)},
		{shape: RightAngledTriangle{3.0, 4.0}, expected: 6.0},
	}

	for _, areaTest := range areaTests {
		checkArea(t, areaTest.shape, areaTest.expected)
	}
}
