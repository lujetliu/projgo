package main

import (
	"chatroom/server"
	"fmt"
	"log"
	"net/http"
)

var (
	addr   = ":2022"
	banner = "Chatroom Start..."
)

func main() {
	fmt.Printf(banner+"\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
