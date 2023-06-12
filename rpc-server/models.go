package main

type Message struct {
	ChatID    string `json:"chat_id"`
	Author    string `json:"author"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
