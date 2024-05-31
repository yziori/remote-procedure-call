package handlers

import (
	"bytes"
	"encoding/json"
	"net"
	"rpc-server/internal/models"
	"testing"
)

type mockConn struct {
	net.Conn
	readBuf  *bytes.Buffer
	writeBuf *bytes.Buffer
}

func (m *mockConn) Read(b []byte) (int, error) {
	return m.readBuf.Read(b)
}

func (m *mockConn) Write(b []byte) (int, error) {
	return m.writeBuf.Write(b)
}

func (m *mockConn) Close() error {
	return nil
}

func TestHandleConnection(t *testing.T) {
	// モックされた接続を作成
	req := models.Request{
		Method:     "floor",
		Params:     []interface{}{10.5},
		ParamTypes: []string{"float64"},
		ID:         1,
	}
	reqBytes, _ := json.Marshal(req)
	readBuf := bytes.NewBuffer(reqBytes)
	writeBuf := &bytes.Buffer{}

	conn := &mockConn{
		readBuf:  readBuf,
		writeBuf: writeBuf,
	}

	server := &RPCServer{}
	server.handleConnection(conn)

	// レスポンスをデコード
	var res models.Response
	if err := json.NewDecoder(writeBuf).Decode(&res); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// 期待されるレスポンスと比較
	if res.Error != "" {
		t.Errorf("Expected no error, but got %s", res.Error)
	}

	expected := 10.0
	if res.Result != expected {
		t.Errorf("Expected result %v, but got %v", expected, res.Result)
	}

	if res.ResultType != "int" {
		t.Errorf("Expected result type int, but got %s", res.ResultType)
	}
}
