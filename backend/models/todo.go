package models

type Todo struct {
	Text   string
	Status bool
}

type UpdateTextRequest struct {
	ID   int
	Text string
}
