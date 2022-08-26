package util

import (
	"fmt"
	"os/exec"
)

// DetectGoBin 是否安装了 Go
func DetectGoBin() {
	result, err := exec.LookPath("go")
	if err != nil {
		panic(err)
	}
	fmt.Print(result)
}

// FindGoMod 查找 go.mod 文件的位置
func FindGoMod() (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}
