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

const (
	GPT_API_URL   = "https://api.openai.com/v1/chat/completions"
	GPT_MODEL     = "gpt-3.5-turbo"
	MESSAGE_SETUP = "You are the Dracula."
	TEMPERATURE   = 0.7
	N             = 1
)

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

	messages := []entities.Message{{Role: "system", Content: MESSAGE_SETUP}, {Role: "user", Content: gptQuery.Prompt}}

	jsonInput, err := json.Marshal(entities.GptRequest{
		Model:       GPT_MODEL,
		Messages:    messages,
		Temperature: TEMPERATURE,
		N:           N,
	})

	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", GPT_API_URL, bytes.NewBuffer(jsonInput))
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
