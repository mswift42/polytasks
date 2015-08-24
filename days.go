package days

import (
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
	Id       int64        `json:"id" datastore:"-"`
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
