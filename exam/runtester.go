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

func RunTester(exercise string, level int, typeCheckpoint string) (bool, int) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	rootDir := filepath.Dir(currentDir)
	pathToTester := filepath.Join(rootDir, ".subject", typeCheckpoint, strconv.Itoa(level), exercise, "main.go")
	ex := exec.Command("go", "run", pathToTester)
	if ex.Err != nil {
		fmt.Println(ex.Err)
		return false, -1
	}
	return true, level + 1
}
