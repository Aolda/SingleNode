package compiler

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteJS(fileName string, funcName string, args []string) string {
	mergedArgs := strings.Join(args, ",")

	cmd := exec.Command("node", "./compiler/script.js", fileName, funcName, mergedArgs)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
