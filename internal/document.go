package internal

type Document struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
}

type Filter struct {
	Key string `json:"key"`

	// omitemptyで省略可能にする
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "ImProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
