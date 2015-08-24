package days

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
