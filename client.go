package main

import (
	"github.com/gorilla/websocket"
)

/*
クライアント・モデル
*/
type client struct {
	socket *websocket.Conn // websocketへのコネクション
	send   chan *task      // メッセージ
	room   *group          // 所属するチャットルーム
}

/*
websocketに書き出されたメッセージを読み込む。
*/
func (c *client) read() {

	// websocketからjson形式でメッセージを読み出し、forwardチャネルに流す。
	// 読み込みは無限ループで実行される。
	for {
		var msg *task
		if err := c.socket.ReadJSON(&msg); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

/*

 */
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
