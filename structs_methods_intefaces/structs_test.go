package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Reactangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Reactangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()

		t.Run(test.name, func(t *testing.T) {
			if got != test.hasArea {
				t.Errorf("%#v got %g want %g", test.shape, got, test.hasArea)
			}
		})
	}

}
