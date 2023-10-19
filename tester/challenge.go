//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package tester

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Program(rootSolutions, exercise string, args ...string) {
	console := func(out string) string {
		var quotedArgs []string
		for _, arg := range args {
			quotedArgs = append(quotedArgs, strconv.Quote(arg))
		}
		s := "\n$ "
		return fmt.Sprintf(s+"go run . %s\n%s$", strings.Join(quotedArgs, " "), out)
	}
	piscineOut, err := GetPiscineAnswer(exercise, args...)
	solutionsOut, err1 := GetSolution(rootSolutions, exercise, args...)
	if err != nil {
		Fatalln("Your program fails (non-zero exit status) when it should not :\n" +
			console(err.Error()) +
			"\n\nExpected :\n" +
			console(solutionsOut))
	}
	if err1 != nil {
		Fatalln("no such file or directory " + err1.Error())
	}
	if piscineOut != solutionsOut {
		Fatalln("Your program output is not correct :\n" +
			console(piscineOut) +
			"\n\nExpected :\n" +
			console(solutionsOut))
	}
}

func GetPiscineAnswer(exercise string, args ...string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	rootDir := filepath.Dir(currentDir)
	pathToPiscineMain := filepath.Join(rootDir, "piscine", exercise, "main.go")
	cmdArgs := append([]string{"run", pathToPiscineMain}, args...)
	ex := exec.Command("go", cmdArgs...)
	out, err := ex.Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			return "", errors.New(string(exitErr.Stderr))
		}
		return "", err
	}
	return string(out), nil
}

func GetSolution(rootSolutions, exercise string, args ...string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	rootDir := filepath.Dir(currentDir)
	pathToPiscineMain := filepath.Join(rootDir, rootSolutions, "main.go")
	cmdArgs := append([]string{"run", pathToPiscineMain}, args...)
	ex := exec.Command("go", cmdArgs...)
	out, err := ex.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
func Fatalln(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
