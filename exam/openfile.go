//  █████╗  ██████╗ █████╗ ██████╗ ███████╗███╗   ███╗██╗███████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔════╝████╗ ████║██║██╔════╝    ██╔═══██╗████╗  ██║██╔════╝
// ███████║██║     ███████║██║  ██║█████╗  ██╔████╔██║██║█████╗      ██║   ██║██╔██╗ ██║█████╗
// ██╔══██║██║     ██╔══██║██║  ██║██╔══╝  ██║╚██╔╝██║██║██╔══╝      ██║   ██║██║╚██╗██║██╔══╝
// ██║  ██║╚██████╗██║  ██║██████╔╝███████╗██║ ╚═╝ ██║██║███████╗    ╚██████╔╝██║ ╚████║███████╗
// ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝     ╚═════╝ ╚═╝  ╚═══╝╚══════╝

package exam

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func OpenFile(idxCheck, level int, task string) (string, string, string) {
	var selectCheckpoint string
	var randomTask string

	switch idxCheck {
	case 1:
		selectCheckpoint = "checkpoint-1"
	case 2:
		selectCheckpoint = "checkpoint-2"
	case 3:
		selectCheckpoint = "checkpoint-3"
	case 4:
		selectCheckpoint = "final-checkpoint"
	}

	tasks := AddTask(selectCheckpoint)

	rand.Seed(time.Now().UnixNano())
	if task != "" {
		randomTask = task
	} else {
		if len(tasks[level]) == 0 {
			panic(errors.New("invalid level"))
		}
		randomTask = tasks[level][rand.Intn(len(tasks[level]))]
	}
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
