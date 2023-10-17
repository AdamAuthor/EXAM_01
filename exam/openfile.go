//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
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
	"math/rand"
	"os"
	"strconv"
	"time"
)

func OpenFile(idxCheck, level int, task string) (string, string, string) {
	var selectCheckpoint string
	tasks := make(map[int][]string)
	tasks[1] = []string{"countdown", "strlen"}

	switch idxCheck {
	case 1:
		fmt.Println(1)
		return "", "", ""
	case 2:
		fmt.Println(2)
		return "", "", ""
	case 3:
		fmt.Println(3)
		return "", "", ""
	case 4:
		selectCheckpoint = "final-checkpoint"
	}

	rand.Seed(time.Now().UnixNano())
	randomTask := tasks[level][rand.Intn(len(tasks[level]))]
	f, err := os.Open("./.subjects/" + selectCheckpoint + "/" + strconv.Itoa(level) + "/" + randomTask + "/attachment/subject.en.txt")
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
	return totalTask, randomTask, selectCheckpoint
}
