package exam

import (
	"os"
)

func CreateDir() {
	if err := os.Mkdir("piscine", os.ModePerm); err != nil {
		return
	}
}
