package exam

import (
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	switch platform := runtime.GOOS; platform {
	case "darwin", "linux":
		cmd := exec.Command("reset")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		// fallback for unsupported OSes: just print many newlines
		cmd := exec.Command("clear")
		cmd.Run()
	}
	os.Exit(0)
}
