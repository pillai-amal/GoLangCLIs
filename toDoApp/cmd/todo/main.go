package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/toDoApp"
)

const toDoFileName = ".todo.json"

func main() {
	l := &toDoApp.List{}

	if err := l.Get(toDoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// printing th list of items when the user has not provided any arguments
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(toDoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
