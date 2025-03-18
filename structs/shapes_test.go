package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		result := shape.Perimeter()
		if result != expected {
			t.Errorf("got %g but expected %g from %#v", result, expected, shape)
		}
	}

	perimeterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 4.0, Height: 10.0}, expected: 28.0},
		{name: "Circle", shape: Circle{Radius: 4.0}, expected: math.Pi * 8},
	}

	for _, pt := range perimeterTests {
		t.Run(pt.name, func(t *testing.T) {
			checkPerimeter(t, pt.shape, pt.expected)
		})

	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("got %g but expected %g", result, expected)
		}
	}

	t.Run("with rectangle", func(t *testing.T) {
		rectangle := Rectangle{4.0, 10.0}
		expected := 40.0
		checkArea(t, rectangle, expected)
	})

	t.Run("with circle", func(t *testing.T) {
		circle := Circle{4.0}
		expected := math.Pi * 16
		checkArea(t, circle, expected)

	})
}
