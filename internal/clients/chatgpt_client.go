package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChatGPTClient struct {
	APIKey      string
	APIEndpoint string
}

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func NewChatGPTClient(apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		APIKey:      apiKey,
		APIEndpoint: "https://api.openai.com/v1/chat/completions",
	}
}

func (c *ChatGPTClient) SendMessage(message string) (string, error) {
	req := ChatGPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "user", Content: message},
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", c.APIEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatGPTResp ChatGPTResponse
	err = json.Unmarshal(body, &chatGPTResp)
	if err != nil {
		return "", err
	}

	if len(chatGPTResp.Choices) > 0 {
		return chatGPTResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from ChatGPT")
}
