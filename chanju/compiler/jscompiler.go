package compiler

import (
	"fmt"
	"os/exec"
	"strings"
)

// func main() {
// 	action := flag.String("action", "", "insert or another")
// 	fileName := flag.String("fileName", "", "what file")
// 	functionName := flag.String("functionName", "", "what function")
// 	args := os.Args[4:]

// 	flag.Parse()

// 	if *action == "insert" {
// 		ExecuteJS(*fileName, *functionName, args)
// 	} else {
// 		fmt.Printf("등록ㄴㄴ\n")
// 	}
// }

func ExecuteJS(fileName string, functionName string, args []string) string {
	// fmt.Printf("%s\n%s\n%s\n", fileName, functionName, args)
	mergedArgs := strings.Join(args, ",")

	cmd := exec.Command("node", fileName, functionName, mergedArgs)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	result := string(out)
	fmt.Println(result)
	return result
}
