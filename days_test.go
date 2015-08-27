package days

import (
	"testing"

	"appengine/aetest"
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

func TestKeyForID(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	key := keyForID(c, 0)
	if key.Kind() != "Task" {
		t.Error("expected <Task>, got: ", key.Kind())
	}
	if key.IntID() != 0 {
		t.Error("expected <0>, got: ", key.IntID())
	}
}
