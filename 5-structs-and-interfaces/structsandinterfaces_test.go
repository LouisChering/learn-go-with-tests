package structsandinterfaces_test

import (
	"testing"

	structsandinterfaces "github.com/louischering/learn-go-with-tests/5-structs-and-interfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := structsandinterfaces.Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   structsandinterfaces.Shape
		hasArea float64
	}{
		{shape: structsandinterfaces.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: structsandinterfaces.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: structsandinterfaces.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
