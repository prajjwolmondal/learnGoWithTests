package structsMethodsAndInterfaces

import "testing"

func TestPermiter(t *testing.T) {
	rectangle := Rectangle{10.0, 2.0}
	got := rectangle.Permiter()
	want := 24.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) // the ".2" means print 2 decimal places
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 20, Height: 20}, hasArea: 400.0},
		{name: "Circle", shape: Circle{Radius: 20}, hasArea: 1256.6370614359173},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
