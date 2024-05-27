package main

import (
	"rpc-server/internal/handlers"
)

// main関数でRPCサーバーを起動
func main() {
	server := &handlers.RPCServer{}
	server.Start()
}
