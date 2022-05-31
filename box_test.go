package golang_united_school_homework

import "testing"

func TestNewBox(t *testing.T) {
	res := NewBox(3)
	if res.shapesCapacity != 3 {
		t.Errorf("Error: only a positive integer is accepted. Expected result 3")
	}
}
func TestAddShapeOverflowBox(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res:=box.AddShape(Rectangle{Height: 3, Weight: 4})
    if  !res == error {
        return fmt.Errorf("Number of figures exceeded. Сurrent value 2")
    }
}
func TestAddShapeInvalidAdd(t *testing.T) {
    box := NewBox(2)
    first_len:= len(box.shapes)
    box.AddShape(Circle{Radius: 1})
    if  first_len == len(box.shapes) {
        return fmt.Errorf("Figure is not added to the list. Expected value: 1")
    }
}
func TestGetByIndexInvalidIndex(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res:=box.GetByIndex(2)
    if  !res == error {
        return fmt.Errorf("Index out of range. Сurrent range: 2")
    }
}
func TestGetByIndexReturnItem(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res,ok:=box.GetByIndex(1)
    if  !ok == nil {
        return fmt.Errorf("Return value error. Expected value: {1 2}")
    }
}
func TestExtractByIndexInvalidIndex(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res:=box.ExtractByIndex(2)
    if  !res == error {
        return fmt.Errorf("Index out of range. Сurrent range: 2")
    }
}
func TestExtractByIndexReturnItem(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res,ok:=box.ExtractByIndex(1)
    if  !ok == nil {
        return fmt.Errorf("Return value error. Expected value: {3 4}")
    }
}
func TestExtractByIndexInvalidDel(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    first_len:= len(box.shapes)
    res,ok:=box.ExtractByIndex(1)
    if  first_len == len(box.shapes) {
        return fmt.Errorf("The figure is not removed from the list. Expected value: 2")
    }
}
func TestReplaceByIndexInvalidIndex(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res, ok := box.ReplaceByIndex(2, Circle{Radius: 1})
    if  !res == error {
        return fmt.Errorf("Index out of range. Сurrent range: 2")
    }
}
func TestReplaceByIndexReturnItem(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res, ok := box.ReplaceByIndex(1, Circle{Radius: 1})
    if  !ok == nil {
        return fmt.Errorf("Return value error. Expected value: {3 4}")
    }
}
func TestReplaceByIndexInvalidDel(t *testing.T) {
    box := NewBox(2)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    first_len:= len(box.shapes)
    res, ok := box.ReplaceByIndex(1, Circle{Radius: 1})
    if  first_len != len(box.shapes) {
        return fmt.Errorf("The figure has not been replaced in the list. Expected value: 2")
    }
}
func TestSumPerimeter(t *testing.T) {
    box := NewBox(3)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Triangle{Side: 2})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res := box.SumPerimeter()
    if  res!= 26.283185307179586 {
        return fmt.Errorf("Incorrect total perimeter of all shapes in the list. Expected value: 26.283185307179586")
    }
}
func TestSumArea(t *testing.T) {
    box := NewBox(3)
    box.AddShape(Circle{Radius: 1})
	box.AddShape(Triangle{Side: 2})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res := box.SumArea()
    if  res!= 16.87364346115867 {
        return fmt.Errorf("Incorrect total area of all shapes in the list. Expected value: 16.87364346115867")
    }
}
func TestRemoveAllCirclesNoCircles(t *testing.T) {
    box := NewBox(2)
	box.AddShape(Triangle{Side: 2})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res := box.RemoveAllCircles()
    if  res== nil {
        return fmt.Errorf("Return value is not correct. The list has no circles")
    }
}
func TestRemoveAllCirclesInvalidDel(t *testing.T) {
    box := NewBox(2)
	box.AddShape(Circle{Radius: 2})
	box.AddShape(Rectangle{Height: 3, Weight: 4})
    res := box.RemoveAllCircles()
    if  res!= nil {
        return fmt.Errorf("Return value is not correct. Expected value: nil")
    }
}
