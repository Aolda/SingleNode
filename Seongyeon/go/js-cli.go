package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	cmd := exec.Command("node", "script.js")
	for _, el := range arr {
		exec.Command(strconv.Itoa(el))
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
