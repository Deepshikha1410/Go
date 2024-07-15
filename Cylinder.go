package Geometry

func CylinderArea(r float64, height float64) float64 {
	return 2 * math.Pi * r * (r + height)
}

func CylinderVolume(r float64, height float64) float64 {
	return math.Pi * r * r * height
}
