package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/toDoApp"
)

const toDoFileName = ".todo.json"

flag.Usage = func () {
    fmt.Fprintf(flag.CommandLine.Output(), "%s tool Developed by AM Pillai @ github.com/pillai_amal", os.Args[0])
    fmt.Fprintf(flag.CommandLine.Output, "contact @ pillai_amal@hotmail.com")
    fmt.Fprintf(flag.CommandLine.Output(), "Usage Deatils")
    flag.PrintDefaults()
}

func main() {
    task := flag.String("task", "", "Task to be included in ToDoList")
    list := flag.Bool("list", false, "List of all tasks")
    complete := flag.Int("complete", 0, "Item to be completed")
    flag.Parse()
	l := &toDoApp.List{}

	if err := l.Get(toDoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// printing th list of items when the user has not provided any arguments
	switch {
    case *list:
		for _, item := range *l {
            if !item.Done {
			    fmt.Println(item.Task)
            }
		}
    case *complete>0:
        if err := l.Complete(*complete); err !=nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
		if err := l.Save(toDoFileName); err != nil{
			fmt.Fprintln(os.Stderr,err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(toDoFileName); err !=nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option provided")
		os.Exit(1)
	}
}
