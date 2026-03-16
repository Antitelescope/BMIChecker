package core

type Healthchecker struct {
}

func NewHealthchecker() *Healthchecker {
	return &Healthchecker{}
}

func (h *Healthchecker) CalculateBMI(height, weight float64) (bmi float64, category string) {
	bmi = weight / (height * height)
	category = getBMICategory(bmi)

	return
}

func getBMICategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "偏瘦"
	case bmi >= 18.5 && bmi < 24:
		return "正常"
	case bmi >= 24 && bmi < 28:
		return "偏旁"
	default:
		return "肥胖"
	}
}
