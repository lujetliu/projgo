package main

import "time"

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

// 给用户发送的消息
type Message struct {
	OwnerID int
	Content string
}

func GetUserId() int {
	return 0
}

func (u User) String() string {
	return ""
}
