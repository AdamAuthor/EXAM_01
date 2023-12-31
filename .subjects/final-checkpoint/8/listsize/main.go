package main

import (
	student "academie/piscine"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func main() {
	table := []solutions.NodeTest{{
		Data: []interface{}{"Hello", "man", "how are you"},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link := &solutions.List{}
		link2 := &List{}
		for _, item := range arg.Data {
			listPushBack(link2, item)
			solutions.ListPushBack(link, item)
		}
		studentLen := solutions.ListSize(link)
		solutionLen := student.ListSize(link2)

		if studentLen != solutionLen {
			challenge.Fatalf("ListSize(%v) == %d instead of %d\n",
				solutions.ListToString(link.Head),
				solutionLen,
				studentLen)
		}
	}
}
