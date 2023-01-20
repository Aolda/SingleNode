package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	var fileName string = "script.js"
	var funcName string = "add"
	var arg1 int = 2
	var arg2 int = 3

	cmd := exec.Command("node", fileName, funcName, strconv.Itoa(arg1), strconv.Itoa(arg2))
	//fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
