package toDoApp_test

import (
	"fmt"
	"github.com/toDoApp"
	"testing"
)

func TestAddFeature(t *testing.T) {
	l := toDoApp.List{}
	taskName := "New Task"
	l.Add(taskName)
	fmt.Printf(l[0].Task)
	if l[0].Task != taskName {
		t.Errorf(" Expected %q and got %q instead", taskName, l[0].Task)
	}
}
