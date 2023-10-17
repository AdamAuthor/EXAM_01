//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package exam

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
)

func RunTester(exercise string, level int, typeCheckpoint string) (int, string) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	rootDir := filepath.Dir(currentDir)
	pathToTester := filepath.Join(rootDir, ".subjects", typeCheckpoint, strconv.Itoa(level), exercise, "main.go")
	ex := exec.Command("go", "run", pathToTester)
	out, err := ex.Output()
	if err != nil || len(out) != 0 {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			// Этот код будет выполнен, если команда завершилась с ошибкой.
			fmt.Println(string(exitErr.Stderr))
			return level, string(exitErr.Stderr)
		} else if !ok {
			// Другой тип ошибки
			return level, err.Error()

		} else {
			fmt.Println(out)
			return level, "out"

		}
	}
	return level + 1, ""
}
