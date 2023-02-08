package main

import (
	"fmt"
	"os/exec"
)

func main() {
	args := []string{"12", "23"}
	a := ExecutePython("add.py", "sum_int", args)
	fmt.Println(a)
}

func ExecutePython(fileName string, funcName string, args []string) string {
	argsForCommand := []string{"script.py", fileName, funcName}
	argsForCommand = append(argsForCommand, args...)
	cmd := exec.Command("/opt/homebrew/bin/python3", argsForCommand...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
