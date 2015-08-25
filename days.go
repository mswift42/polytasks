package days

import (
	"encoding/json"
	"io"

	"appengine"
	"appengine/datastore"
)

type TaskNote struct {
	NoteContent string
}

type TaskCategory struct {
	Category string
}

type Task struct {
	ID       int64        `json:"id" datastore:"-"`
	Summary  string       `json:"summary"`
	Content  TaskNote     `json:"content"`
	Done     bool         `json:"done"`
	Category TaskCategory `json:"category"`
}

func keyForID(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "Task", "", id, tasklistkey(c))
}

func tasklistkey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Task", "default_tasklist", 0, nil)
}

func (t *Task) key(c appengine.Context) *datastore.Key {
	if t.ID == 0 {
		return datastore.NewIncompleteKey(c, "Task", tasklistkey(c))
	}
	return datastore.NewKey(c, "Task", "", t.ID, tasklistkey(c))
}

func (t *Task) save(c appengine.Context) (*Task, error) {
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}
	t.ID = k.IntID()
	return t, nil
}

func decodeTask(r io.ReadCloser) (*Task, error) {
	defer r.Close()
	var task Task
	err := json.NewDecoder(r).Decode(&task)
	return &task, err
}
