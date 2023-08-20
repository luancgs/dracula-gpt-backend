package entities

type GptQuery struct {
	Prompt string
}

type GptRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	N           int64     `json:"n"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []choice `json:"choices"`
	Usage   usage    `json:"usage"`
}

type choice struct {
	Index   int64   `json:"index"`
	Message message `json:"message"`
}

type message struct {
	Role          string `json:"role"`
	Content       string `json:"content"`
	Finish_reason string `json:"finish_reason"`
}

type usage struct {
	Prompt_tokens     int64 `json:"prompt_tokens"`
	Completion_tokens int64 `json:"completion_tokens"`
	Total_tokens      int64 `json:"total_tokens"`
}
