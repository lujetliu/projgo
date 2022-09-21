package logic

import "log"

// broadcaster 广播器
type broadcaster struct {
	// 所有聊天室用户
	users map[string]*User

	// 所有 channel 统一管理，可以避免外部乱用
	enteringChannel chan *User
	leavingChannel  chan *User
	messageChannel  chan *Message

	// 判断该昵称用户是否可进入聊天室（重复与否）：true 能，false 不能
	checkUserChannel      chan string
	checkUserCanInChannel chan bool
}

func (b *broadcaster) CanEnterRoom(nickname string) bool {
	b.checkUserChannel <- nickname

	return <-b.checkUserCanInChannel
}

var MessageQueueLen = 100

// 广播器全局应该只有一个，所以是典型的单例
// 使用饿汉式实现广播器的单例模式
var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan *Message, MessageQueueLen),

	checkUserChannel:      make(chan string),
	checkUserCanInChannel: make(chan bool),
}

// Start 启动广播器
// 需要在一个新 goroutine 中运行, 因为它不会返回
func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			// 新用户进入
			b.users[user.NickName] = user
			b.sendUserList()
		case user := <-b.leavingChannel:
			// 用户离开
			delete(b.users, user.NickName)
			// 避免 goroutine 泄露
			user.CloseMessageChannel()
			b.sendUserList()
		case msg := <-b.messageChannel:
			if msg.To == "" {
				// 给所有在线用户发送消息
				for _, user := range b.users {
					if user.UID == msg.User.UID {
						continue
					}
					user.MessageChannel <- msg
				}
			} else {
				if user, ok := b.users[msg.To]; ok {
					user.MessageChannel <- msg
				} else {
					// 对方不在线或用户不存在，直接忽略消息
					log.Println("user:", msg.To, "not exists!")
				}
			}

		case nickname := <-b.checkUserChannel:
			if _, ok := b.users[nickname]; ok {
				b.checkUserCanInChannel <- false
			} else {
				b.checkUserCanInChannel <- true
			}
		}
	}
}

func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

func (b *broadcaster) Broadcast(msg *Message) {
	b.messageChannel <- msg
}

func (b *broadcaster) sendUserList() {
}
