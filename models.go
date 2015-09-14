package main

type ChatMessage struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{}
}
