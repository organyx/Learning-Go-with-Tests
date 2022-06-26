package perimeter

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
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	got := rectangle.Area()
	// 	want := 100.0

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := circle.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// })

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("%#v got %g want %g", shape, got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10, 12}
	// 	checkArea(t, rectangle, 120.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{

		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
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
