package exam

import (
	"log"
	"os"
)

func CreateDir() {
	if err := os.Mkdir("piscine", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
