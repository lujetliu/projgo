package server

import "net/http"

func RegisterHandle() {
	inferRootDir()

	// go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	// http.HandleFunc("/ws", WebSocketHandleFunc)
}
