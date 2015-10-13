package days

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
func TestTaskListKey(t *testing.T) {
	assert := assert.New(t)
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	key := tasklistkey(c)
	assert.Equal(key.StringID(), "default_tasklist")
	assert.Equal(key.Kind(), "Task")
	assert.Equal(key.IntID(), 0)
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

func TestSave(t *testing.T) {
	assert := assert.New(t)
	t1 := Task{Summary: "task1",
		Content: []string{"Some Content"}}
	c, err := aetest.NewContext(nil)
	defer c.Close()
	if err != nil {
		t.Fatal(err)
	}
	tn, err := t1.save(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(tn.Summary, "task1")

}

func TestListTasks(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	assert := assert.New(t)
	t1 := Task{Summary: "task1", Done: true}
	t2 := Task{Summary: "task2", Done: false}
	t1n, err := t1.save(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t1n.Done, true)
	t2n, err := t2.save(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t2n.Done, false)
	tasklist, err := listTasks(c)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(tasklist[0].Summary, "task2")
	assert.Equal(tasklist[1].Done, true)
}

func TestHandlers(t *testing.T) {
	_, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	uri := "/api/tasks/"
	var testjson1 = `{"summary" : "task1",
                "content" : ["taskcontent1"],
    "done": false}`
	instance, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}
	req, err := instance.NewRequest("POST", uri, ioutil.NopCloser(strings.NewReader(testjson1)))
	http.DefaultServeMux.ServeHTTP(resp, req)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()

	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s: ", p)
		}
		if !strings.Contains(string(p), "task1") {
			t.Errorf("header response doesn't match: \n%s", p)
		}
	}

	getreq, err := instance.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, getreq)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return erros: %s: ", p)
		}
		if !strings.Contains(string(p), "task1") {
			t.Errorf("header response doesn't match: \n%s", p)
		}
		if !strings.Contains(string(p), "taskcontent1") {
			t.Errorf("header reponse doesn't match: \n%s", p)
		}
	}
	uri = "/api/task/5629499534213120"
	delreq, err := instance.NewRequest("DELETE", uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(resp, delreq)
	if p, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if strings.Contains(string(p), "Error") {
			t.Errorf("header response shouldn't return error: %s: ", p)
		}
	}
}
