package structs_methods_interfaces
import "testing"

func TestPerimeter(t *testing.T){
	t.Run("Rectangle perimeter", func(t *testing.T) {
		r := Rectangle{Width:  10,Height: 10}
		got := r.Perimeter()
		want := 40.0

		if got != want{
			t.Errorf("I wanted %.2f but I got %.2f", want, got)
		}
	})
}

func TestArea(t *testing.T){
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

