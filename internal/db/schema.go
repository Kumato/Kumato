package db

type User struct {
	Qid   string `json:"qid,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
