package days

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"appengine/aetest"
)

func TestTaskStruct(t *testing.T) {
	t1 := Task{ID: 1, Summary: "task1",
		Content:  []string{"TaskContent1"},
		Done:     false,
		Category: []string{"work", "travel"},
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

func TestDecodeTask(t *testing.T) {
	assert := assert.New(t)
	var testjson1 = `{"summary" : "task1",
                "content" : ["taskcontent1"],
    "done": false}`
	t1, err := decodeTask(ioutil.NopCloser(strings.NewReader(testjson1)))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t1.Summary, "task1")
}
