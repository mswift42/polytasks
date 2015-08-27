package days

import (
	"encoding/json"
	"io"
	"time"

	"appengine"
	"appengine/datastore"
)

// Task represents a single Task object, with an ID, a task summary
// the task content, consisting of a slice of strings, a scheduled
// date, the task status `done` (true or false) and a slice
// of taskCategories.
type Task struct {
	ID        int64     `json:"id" datastore:"-"`
	Summary   string    `json:"summary"`
	Content   []string  `json:"content"`
	Scheduled time.Time `json:"scheduled"`
	Done      bool      `json:"done"`
	Category  []string  `json:"category"`
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

func listTasks(c appengine.Context) (*[]Task, error) {
	tasks := []Task{}
	keys, err := datastore.NewQuery("Task").Ancestor(tasklistkey(c)).Order("-Done").Order("Scheduled").GetAll(c, &tasks)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tasks); i++ {
		tasks[i].ID = keys[i].IntID()
	}
	return &tasks, err
}
