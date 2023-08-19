package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofor-little/env"
	"github.com/luancgs/dracula-gpt-backend/src/entities"
)

const GPT_API_URL = "https://api.openai.com/v1/chat/completions"

type GptService interface {
	CreateQuery(entities.GptQuery) (string, error)
}

type gptService struct{}

func NewGpt() GptService {
	return &gptService{}
}

func (s *gptService) CreateQuery(gptQuery entities.GptQuery) (string, error) {

	apiKey, err := env.MustGet("GPT_API_KEY")
	if err != nil {
		fmt.Println("Error getting API key:", err)
		return "", err
	}

	jsonInput := fmt.Sprintf(`{
		"model": "gpt-3.5-turbo",
		"messages": [{"role": "user", "content": "I want you to speak with me as if you were the Dracula!"}, {"role": "assistant", "content": "Very well. I shall embrace the darkness and speak to you in the voice of Dracula."}, {"role": "user", "content": "%s"}],
		"temperature": 0.7,
		"n": 1
	}`, gptQuery.Prompt)

	postBody := []byte(jsonInput)

	req, err := http.NewRequest("POST", GPT_API_URL, bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	response := entities.GptResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
