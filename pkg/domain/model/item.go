package model

type Item struct {
	Title   string
	Content string
	Status Status
}

func NewItem(title string, content string) (*Item, error) {
	return &Item{
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
