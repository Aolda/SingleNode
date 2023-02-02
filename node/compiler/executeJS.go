package compiler

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteJS(fileName string, funcName string, args []string) string {
	mergedArgs := strings.Join(args, ",")

	cmd := exec.Command("node", "./compiler/script.js", fileName, funcName, mergedArgs)
	fmt.Println(cmd)
	fmt.Println(fileName == "myScript.js")
	fmt.Println(funcName == "add")
	fmt.Println(mergedArgs == "1 2")
	fmt.Println(cmd.String() == "/Users/odongjae/.nvm/versions/node/v18.12.1/bin/node ./compiler/script.js myScript.js add 1 2")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
	fmt.Println(string(out))
	return string(out)
}
