package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTable := []struct {
		name  string
		shape Shape
		hasArea  float64
	}{
    {name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.159253589793},
		{name: "Triangle", shape: Triangle{Base: 6, Height: 12}, hasArea: 36.0},
	}

	for _, tt := range areaTable {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}