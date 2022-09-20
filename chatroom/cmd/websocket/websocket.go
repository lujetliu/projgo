package main

/*
	WebSocket 是一种与 HTTP 不同的协议, 两者都位于 OSI 模型的应用层,
	并且都依赖于传输层的 TCP 协议;
	RFC 6455 规定: "WebSocket 设计为通过 80 和 443 端口工作, 以及支持 HTTP
	代理和中介", 从而使其与 HTTP 协议兼容;  为了实现兼容性 ,WebSocket 握手使用
	HTTP Upgrade 头从 HTTP 协议更改为 WebSocket 协议;

	WebSocket 协议支持 Web 浏览器(或其他客户端应用程序)与 Web 服务器之间的交互,
	具有较低的开销, 便于实现客户端与服务器的实时数据传输; 服务器可以通过标准化
	的方式来实现, 而无需客户端首先请求内容, 并允许消息在保持连接打开的同时来回
	传递; 通过这种方式可以在客户端和服务器之间进行双向持续交互;
	通信默认通过 TCP 端口 80 或 443 完成;

	WebSocket 提供全双工通信, WebSocket 还可以在 TCP 之上启用消息流;
	TCP 单独处理字节流, 没有固有的消息概念; 在 WebSocket 之前, 使用
	Comet 可以实现全双工通信,  但是 Comet 存在 TCP 握手和 HTTP 头的开销,
	因此对于小消息来说效率很低, WebSocket 协议旨在解决这些问题;

	Websocket 通过 HTTP/1.1 协议的 101 状态码进行握手;

	TODO: websocket 协议, 使用本节的 ./server.go 和 ./client.go 进行通信, 并
	使用 wireshark 抓包熟悉协议细节

*/
