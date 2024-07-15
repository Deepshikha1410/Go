package Geometry

func Area() float64 {
	return r.width * r.height
}

func Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func Volume() float64 {
	return r.width * r.height * r.depth
}