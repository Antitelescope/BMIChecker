package main

import (
	"fmt"
	"healthchecker/core"
	"healthchecker/input"
	"healthchecker/output"
)

func main() {
	inputer := input.TerminalInputer{}
	outputer := output.TerminalOutputer{}
	checker := core.NewHealthchecker()

	for {
		height, weight, err := inputer.GetBodyInfo()
		if err != nil {
			fmt.Println("输入有误：", err)
			continue
		}

		bmi, category := checker.CalculateBMI(height, weight)
		outputer.OutputResult(bmi, category)

		if !askContinue() {
			break
		}
	}

}

func askContinue() bool {
	fmt.Print("是否继续?(y/n):")
	var again string
	_, scan := fmt.Scan(&again)
	if scan != nil {
		return false
	}
	return again == "y" || again == "Y"
}
