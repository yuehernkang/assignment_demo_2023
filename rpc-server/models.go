package main

type Message struct {
	Chat     string `json:"chat"`
	Text     string `json:"text"`
	Sender   string `json:"sender"`
	SendTime int64  `json:"sendtime"`
}
