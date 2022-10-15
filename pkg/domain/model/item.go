package model

type Item struct {
	Title   string
	Content string
}

func NewItem(title string, content string) (*Item, error) {
	return &Item{
		Title:   title,
		Content: content,
	}, nil
}
