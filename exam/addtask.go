package exam

import (
	"fmt"
	"os"
	"strconv"
)

func AddTask(checkpoint string) (map[int][]string, int) {
	tasks := make(map[int][]string)
	files, err := os.ReadDir("./.subjects/" + checkpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, 0
	}
	var exams []string

	for i := range files {
		taskNames, err := os.ReadDir("./.subjects/" + checkpoint + "/" + strconv.Itoa(i+1) + "/")
		if err != nil {
			fmt.Println("Error:", err)
			return nil, 0
		}
		for _, name := range taskNames {
			exams = append(exams, string(name.Name()))
		}
		tasks[i+1] = exams
		exams = []string{}
	}

	return tasks, len(tasks)
}

func (c *Checkpoint) GetByCheckpoint() int {
	var lenTask int
	switch c.IdxCheck {
	case 1:
		_, lenTask = AddTask("checkpoint-1")
	case 2:
		_, lenTask = AddTask("checkpoint-2")
	case 3:
		_, lenTask = AddTask("checkpoint-3")
	case 4:
		_, lenTask = AddTask("final-checkpoint")
	default:
		return 0
	}
	return lenTask
}
