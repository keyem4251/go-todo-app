package model

type Item struct {
	Id int
	Title   string
	Content string
	Status Status
}

func NewItem(id int, title string, content string) (*Item, error) {
	return &Item{
		Id: id,
		Title:   title,
		Content: content,
		Status: Open,
	}, nil
}

type Status int

const (
	Open Status = iota
	InProgress
	Done
)
