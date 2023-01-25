package main

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()

	// vm.Set("addFunc", func(call otto.FunctionCall) int {
	// 	args1, _ := call.Argument(0).ToString()
	// 	args2, _ := call.Argument(1).ToString()
	// 	int1, _ := strconv.Atoi(args1)
	// 	int2, _ := strconv.Atoi(args2)
	// 	return int1 + int2
	// })
	// value, _ := vm.Call("addFunc", nil, 1, 2)
	// fmt.Println(value)

	file, _ := ioutil.ReadFile("useObject.js")

	fileString := string(file)

	vm.Set("fnName", "mul")
	vm.Set("inputArgs", "1,2")
	value, err := vm.Run(fileString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
