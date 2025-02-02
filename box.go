package golang_united_school_homework

import (
	"fmt"
)

const InvalidIndex string = "figure at index does not exist or index is out of range"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return fmt.Errorf("quantity cannot be exceeded, current value: %d", len(b.shapes))
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf(InvalidIndex)
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf(InvalidIndex)
	}
	res := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return res, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("figure at index does not exist or index is out of range")
	}
	res := b.shapes[i]
	b.shapes[i] = shape
	return res, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64
	for _, item := range b.shapes {
		res = res + item.CalcPerimeter()
	}
	return res

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64
	for _, item := range b.shapes {
		res = res + item.CalcArea()
	}
	return res

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	noCircle := []Shape{}
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			noCircle = append(noCircle, shape)
		}
	}
	if len(noCircle) == len(b.shapes) {
		return fmt.Errorf("there are no circles in the list")
	}
	b.shapes = noCircle
	return nil

}
