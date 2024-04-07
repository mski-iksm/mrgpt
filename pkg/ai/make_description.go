package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

func MakeDescription(gitDiff string, apiKey string) string {

	// チャットリクエストを作成
	chatReq := ChatRequest{
		Model: "gpt-4-0125-preview",
		Messages: []Message{
			{
				Role: "system",
				Content: "You are an excellent engineer. " +
					"Please create a pull request description in Japanese for the git diff " +
					"I'm about to give you." +
					"However, the format should be markdown and the following items should be chaptered as H2 tags." +
					"概要," +
					"本MR前の問題点," +
					"本MRでの対応方法,",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("git diff: %s", gitDiff),
			},
		},
	}

	// リクエストボディをJSONに変換
	requestBody, err := json.Marshal(chatReq)
	if err != nil {
		panic(err)
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// HTTPクライアントを作成してリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// レスポンスを読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// レスポンスをパース
	var chatResp ChatResponse
	err = json.Unmarshal(body, &chatResp)
	if err != nil {
		panic(err)
	}

	return chatResp.Choices[0].Message.Content
}
