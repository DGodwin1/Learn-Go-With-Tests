package structs_methods_interfaces
import "testing"

func TestPerimeter(t *testing.T){
	t.Run("test perimeter is correct", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want{
			t.Errorf("I wanted %.2f but I got %.2f", want, got)
		}
	})
}

func TestArea(t *testing.T){
	t.Run("Test that area is correct", func(t *testing.T){
		got := Area(10,20)
		want := 200.0

		if got != want{
			t.Errorf("I wanted %.2f but I got %.2f", want, got)
		}
	})
}