package golang_united_school_homework

import "testing"

func TestNewBox(t *testing.T) {
	res := NewBox(3)
	if res.shapesCapacity != 3 {
		t.Errorf("Error: ")
	}

}
