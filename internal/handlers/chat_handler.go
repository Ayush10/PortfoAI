package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ayush10/PortfoAI/internal/clients"
)

type ChatHandler struct {
	chatGPTClient *clients.ChatGPTClient
}

func NewChatHandler(apiKey string) *ChatHandler {
	return &ChatHandler{
		chatGPTClient: clients.NewChatGPTClient(apiKey),
	}
}

func (h *ChatHandler) HandleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Message string `json:"message"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	response, err := h.chatGPTClient.SendMessage(req.Message)
	if err != nil {
		http.Error(w, "Error getting AI response", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"response": response})
}
