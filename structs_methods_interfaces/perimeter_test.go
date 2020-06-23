package structs_methods_interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle perimeter", func(t *testing.T) {
		r := Rectangle{Width: 10, Height: 10}
		got := r.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("I wanted %.2f but I got %.2f", want, got)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: Rectangle{Width: 12, Height: 6},
			want:  72.0},

		{
			name:  "Circle",
			shape: Circle{radius: 10},
			want:  314.1592653589793},

		{
			name:  "Triangle",
			shape: Triangle{Base: 10, Height: 10},
			want:  50.0},
	}

	// Iterate over each item in the areaTest slice.
	// We don't care about the index, which is why we use _
	for _, tt := range areaTests {

		// At each item in the slice, look at the shape (.shape) and then call its Area() method
		got := tt.shape.Area()

		// See if the value for the shape's area is the same as the corresponding 'want'
		// value for that entry in the slice.
		if got != tt.want {
			t.Errorf("Area received: %g. Area wanted: %g. Tested: %v", got, tt.want, tt.name)
		}
	}
}

//	////checkArea := func(t *testing.T, shape Shape, want float64) {
//	////	t.Helper()
//	////	got := shape.Area()
//	////	if got != want {
//	////		t.Errorf("got %g want %g", got, want)
//	////	}
//	////}
//	//
//	//t.Run("rectangles", func(t *testing.T) {
//	//	rectangle := Rectangle{12, 6}
//	//	checkArea(t, rectangle, 72.0)
//	//})
//	//
//	//t.Run("circles", func(t *testing.T) {
//	//	circle := Circle{10}
//	//	checkArea(t, circle, 314.1592653589793)
//	//})
//	//
//}
