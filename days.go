package days

type TaskNote struct {
	NoteContent string
}

type Task struct {
	Id      int64    `json:"id" datastore:"-"`
	Summary string   `json:"summary"`
	Content TaskNote `json:"content"`
}
