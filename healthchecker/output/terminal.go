package output

import "fmt"

type ResultOutputer interface {
	OutputResult(bmi float64, category string)
}

type TerminalOutputer struct{}

func (t TerminalOutputer) OutputResult(bmi float64, category string) {
	fmt.Printf("测试BMI: %.2f\t评级为%s\n", bmi, category)
}
