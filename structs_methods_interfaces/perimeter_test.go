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
	t.Run("Rectangle area", func(t *testing.T){
		r := Rectangle{Width: 10, Height: 20}
		got := r.Area()
		want := 200.0

		if got != want{
			t.Errorf("I wanted %.2f but I got %.2f", want, got)
		}
	})

	t.Run("Circle area", func(t *testing.T){
		c := Circle{radius: 10}
		got := c.Area()
		want := 314.1592653589793

		if got != want{
			t.Errorf("I wanted %g but I got %g", want, got)
		}
	})
}

