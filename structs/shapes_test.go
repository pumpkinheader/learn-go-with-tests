package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{20.0, 30.0}, want: 600},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{20.0, 30.0}, want: 300.0},
	}
	for _, tt := range areaTests{
		t.Run(tt.name, func(t *testing.T){
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}