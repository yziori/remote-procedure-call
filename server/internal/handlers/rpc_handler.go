package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"rpc-server/internal/models"
	"rpc-server/internal/services"
)

// RPCサーバーの構造体
type RPCServer struct {
	listener net.Listener
}

// Start関数でRPCサーバーを起動
func (server *RPCServer) Start() {
	var error error
	server.listener, error = net.Listen("unix", "/tmp/unix_socket")
	if error != nil {
		log.Fatal("Error starting RPC server: ", error)
	}
	defer server.listener.Close()
	fmt.Println("Server is listening on /tmp/unix_socket")

	// クライアントからの接続を待ち続ける
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection: ", err)
			continue
		}
		// 新しい接続毎に新しいgoroutineを起動
		go server.handleConnection(conn)
	}
}

// handleConnection関数でリクエストを処理
func (server *RPCServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	var req models.Request
	decoder := json.NewDecoder(conn)
	// リクエストをデコード
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error decoding request: ", err)
		return
	}

	var res models.Response
	res.ID = req.ID

	// リクエストのメソッドによって処理を分岐
	switch req.Method {
	case "floor":
		res = services.Floor(req)
	}
}
