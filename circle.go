package golang_united_school_homework

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) CalcPerimeter()float64{
	pi := math.Pi
	return 2*pi*c.Radius
}

func (c Circle) CalcArea()float64{
	pi := math.Pi
	return pi*c.Radius*c.Radius
}
