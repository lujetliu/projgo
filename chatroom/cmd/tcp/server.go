package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	// 新用户到来, 通过该 channel 进行登记
	enteringChannel = make(chan *User)
	// 用户离开, 通过该 channel 进行注销
	leavingChannel = make(chan *User)
	// 广播专用的用户普通消息 channel, 缓冲是尽可能避免出现异常情况堵塞,
	// 消息通道的长度具体值根据情况调整
	messageChannel = make(chan Message, 8)
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1. 新用户进来，构建该用户的实例
	user := &User{
		ID:             GetUserId(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 2. 当前在一个新的 goroutine 中，用来进行读操作，因此需要开一个 goroutine 用于写操作
	// 读写 goroutine 之间可以通过 channel 进行通信
	go sendMessage(conn, user.MessageChannel)

	// 3. 给当前用户发送欢迎信息；给所有用户告知新用户到来
	user.MessageChannel <- "Wecome, " + user.String()
	messageChannel <- Message{
		OwnerID: user.ID,
		Content: "user:`" + strconv.Itoa(user.ID) + "` has enter",
	}

	// 4. 将该记录到全局的用户列表中，避免用锁
	enteringChannel <- user

	// 自动踢出未活跃用户(5分钟内未发送消息)
	var userActive = make(chan struct{})
	go func() {
		d := 5 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close() // 断开连接, 用户离开
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	// 5. 循环读取用户的输入
	// Scanner 处理如按行读取输入序列或空格分隔单词等
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- Message{
			OwnerID: user.ID,
			Content: strconv.Itoa(user.ID) + ":" + input.Text(),
		}

		// 用户活跃
		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)
	}

	// 6. 用户离开
	leavingChannel <- user
	messageChannel <- Message{
		OwnerID: user.ID,
		Content: "user:`" + strconv.Itoa(user.ID) + "` has left",
	}
}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// broadcaster 用于记录聊天室用户, 并进行消息广播:
// 1. 新用户进来
// 2. 用户普通消息
// 3. 用户离开
func broadcaster() {
	users := make(map[*User]struct{})

	for {
		select {
		case user := <-enteringChannel:
			// 新用户进入
			users[user] = struct{}{}
		case user := <-leavingChannel:
			// 用户离开
			delete(users, user)
			// 避免 goroutine 泄露
			close(user.MessageChannel)
		case msg := <-messageChannel: // 广播消息
			// 给所有在线用户发送消息
			for user := range users {
				if user.ID == msg.OwnerID {
					// 自己的消息不发给自己
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}
