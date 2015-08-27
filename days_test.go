package days

import (
	"testing"
)

func TestTaskStruct(t *testing.T) {
	tn1 := TaskNote{NoteContent: "NoteContent1"}
	cats := []TaskCategory{
		TaskCategory{Category: "work"},
		TaskCategory{Category: "travel"},
	}
	t1 := Task{ID: 1, Summary: "task1",
		Content:  tn1,
		Done:     false,
		Category: cats,
	}
	if t1.Summary != "task1" {
		t.Error("expected <task1> got: ", t1.Summary)
	}
}
