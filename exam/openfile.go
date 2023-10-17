// █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package exam

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenFile(idx int) string {
	var selectCheckpoint string
	switch idx {
	case 1:
		fmt.Println(1)
		return ""
	case 2:
		fmt.Println(2)
		return ""
	case 3:
		fmt.Println(3)
		return ""
	case 4:
		selectCheckpoint = "final-checkpoint"
	}

	

	f, err := os.Open("./.subjects/" + selectCheckpoint + "/1/countdown/attachment/subject.en.txt")
	if err != nil {
		panic(err)
	}

	// Create a buffered reader for the file.
	r := bufio.NewReader(f)
	var totalTask string
	// Read the contents of the file line by line.
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		totalTask += line
	}

	// Close the file.
	f.Close()
	return totalTask
}
