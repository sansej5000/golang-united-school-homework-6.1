package golang_united_school_homework

import "testing"

func TestNewBox(t *testing.T) {
	res := NewBox(3)
	if res.shapesCapacity != 3 {
		t.Errorf("Error: only a positive integer is accepted. Expected result 3")
	}
}

func TestAddShape(t *testing.T) {
	box := NewBox(3)
	figure := Circle{Radius: 2}
	ok := box.AddShape(figure)
	if ok != nil {
		t.Errorf("Error: only a positive integer is accepted")
	}
}
