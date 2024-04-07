package ai

// リクエストボディの構造体
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// メッセージ構造体
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// APIレスポンスの構造体
type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}
