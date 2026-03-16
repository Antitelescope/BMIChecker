package input

import (
	"errors"
	"fmt"
)

// BodyInfoInputer 输入器接口
type BodyInfoInputer interface {
	GetBodyInfo() (height, weight float64, err error)
}

// TerminalInput 终端实现
type TerminalInputer struct{}

func (t TerminalInputer) GetBodyInfo() (height, weight float64, err error) {
	height, err1 := getValidInput("请输入身高（米）")
	weight, err2 := getValidInput("请输入体重（公斤）")

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("输入无效")
	}

	return height, weight, nil
}

func getValidInput(prompt string) (float64, error) {
	fmt.Println(prompt)
	var value float64
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, errors.New("输入格式错误，请输入数字")
	}

	if value <= 0 {
		return 0, errors.New("输入必须大于0")
	}

	return value, nil
}
