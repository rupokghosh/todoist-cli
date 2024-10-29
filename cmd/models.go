package cmd

type Task struct {
	Content  string `json:"content"`
	Due      string `json:"due_string"`
	Priority int    `json:"priority"`
	DueLang  string `json:"due_lang"`
}
