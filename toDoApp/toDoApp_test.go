package toDoApp_test

import (
	"fmt"
	"github.com/toDoApp"
	"io/ioutil"
	"os"
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

func TestCompleteFeature(t *testing.T) {
	l := toDoApp.List{}
	taskName := "New Task"

	l.Add(taskName) //adding tak to list
	if l[0].Task != taskName {
		t.Errorf(" Expected %q but got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("Task is showing as completed which should be")
	}

	l.Complete(1) //completing task
	if !l[0].Done {
		t.Errorf("The should be complete but it showing as incomplete")
	}
}

func TestDeleteFeature(t *testing.T) {
	l := toDoApp.List{}
	taskNames := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}
	for _, v := range taskNames {
		l.Add(v)
	}

	if l[0].Task != taskNames[0] {
		t.Errorf(" Expected %q but got %q instead", taskNames[0], l[0].Task)
	}

	l.Delete(2) //delete operation test
	if len(l) != 2 {
		t.Errorf("Expected %d but got %q", 2, len(l))
	}
}

func TestSavenGet(t *testing.T) {
	l1 := toDoApp.List{}
	l2 := toDoApp.List{}
	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf(" Expected %q but got %q instead", taskName, l1[0].Task)
	}

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error in creating a temporary file %s", err)
	}

	defer os.Remove(tempFile.Name())
	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error occcured during saving file %s", err)
	}
	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error occured during reading file %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Mismatch in written task : %s and read task : %s ", l1[0].Task, l2[0].Task)
	}
}
