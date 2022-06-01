package golang_united_school_homework

type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter()float64{
	return 2*(r.Height+r.Weight)
}

func (r Rectangle) CalcArea()float64{
	return r.Height*r.Weight
}

