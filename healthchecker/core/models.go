package core

type BodyMeasure struct {
	Height   float64
	Weight   float64
	BMI      float64
	Category string
}

func (b *BodyMeasure) Calculate() {
	b.BMI = b.Weight / (b.Height * b.Height)
	b.Category = getBMICategory(b.BMI)
}
