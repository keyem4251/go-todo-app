package model

type Item struct {
	Id      int
	Title   string
	Content string
	Status  Status
}

func NewItem(id int, title string, content string) *Item {
	return &Item{
		Id:      id,
		Title:   title,
		Content: content,
		Status:  Open,
	}
}

type Status int

const (
	Open Status = iota
	InProgress
	Done
)
