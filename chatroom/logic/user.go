package logic

import (
	"context"
	"errors"
	"regexp"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// 系统用户，代表是系统主动发送的消息
var System = &User{UID: 0}

const (
	MsgTypeNormal   = iota // 普通 用户消息
	MsgTypeSystem          // 系统消息
	MsgTypeError           // 错误消息
	MsgTypeUserList        // 发送当前用户列表
)

// 给用户发送的消息
type Message struct {
	// 哪个用户发送的消息
	User    *User            `json:"user"`
	Type    int              `json:"type"`
	Content string           `json:"content"`
	MsgTime time.Time        `json:"msg_time"`
	Users   map[string]*User `json:"users"`

	// 消息发送给谁，表明这是一条私信
	To string `json:"to"`

	// 消息@了谁
	Ats []string `json:"ats"`
}

func NewMessage(u *User, content string) *Message {
	return &Message{
		User:    u,
		Content: content,
		Type:    MsgTypeNormal,
		MsgTime: time.Now(),
		Users:   Broadcaster.users,
	}
}

type User struct {
	UID            int64         `json:"uid"`
	NickName       string        `json:"nickname"`
	EnterAt        time.Time     `json:"enter_at"`
	Addr           string        `json:"addr"`
	MessageChannel chan *Message `json:"-"`

	conn *websocket.Conn
}

func (u *User) SendMessage(ctx context.Context) {
	for msg := range u.MessageChannel {
		wsjson.Write(ctx, u.conn, msg)
	}
}

func (u *User) ReceiveMessage(ctx context.Context) error {
	var (
		receiveMsg map[string]string
		err        error
	)
	for {
		err = wsjson.Read(ctx, u.conn, &receiveMsg)
		if err != nil {
			// 判定连接是否关闭了，正常关闭，不认为是错误
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) { // TODO: 熟悉 errors 的常用函数
				return nil
			}

			return err
		}

		// 内容发送到聊天室
		sendMsg := NewMessage(u, receiveMsg["content"])

		// // 解析 content，看是否是一条私信消息
		// sendMsg.Content = strings.TrimSpace(sendMsg.Content)
		// if strings.HasPrefix(sendMsg.Content, "@") {
		// 	sendMsg.To = strings.SplitN(sendMsg.Content, " ", 2)[0][1:]
		// }
		// 解析 content，看看 @ 谁了
		reg := regexp.MustCompile(`@[^\s@]{4,20}`) // \s 匹配空格
		// @[^\s@]{2,20} 即以@开头连续配置除空格和@之外的4-20个字符
		sendMsg.Ats = reg.FindAllString(sendMsg.Content, -1)

		Broadcaster.Broadcast(sendMsg)
	}
}

func (u *User) CloseMessageChannel() {
	close(u.MessageChannel)
}

func NewNoticeMessage(msg string) *Message {
	message := Message{
		User:    System,
		Type:    MsgTypeSystem,
		Content: msg,
		Users:   Broadcaster.users,
	}
	return &message
}

func NewWelcomeMessage(nickname string) *Message {
	message := Message{
		User:    System,
		Type:    MsgTypeSystem,
		Content: "Welcome, " + nickname,
	}
	return &message
}

func NewUser(conn *websocket.Conn, nickname, addr string) *User {
	return &User{
		UID:            time.Now().UnixNano(),
		NickName:       nickname,
		EnterAt:        time.Now(),
		Addr:           addr,
		MessageChannel: make(chan *Message, 10),
		conn:           conn,
	}
}

func NewErrorMessage(msg string) *Message {
	return &Message{
		Content: msg,
		Type:    MsgTypeError,
		User:    System,
	}
}
