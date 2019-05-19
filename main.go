package main

import (
	"log"
	"net/http"
)

func main() {

	/* ルートへのアクセスに対してハンドラを貼り、group.htmlをサーブする */
	http.Handle("/", &templateHandler{filename: "/group.html"})

	group := newRoom()

	http.Handle("/room", group)

	/* チャットルームを起動する */
	go group.run()

	/* webサーバを開始する */
	log.Println("webサーバを開始します。")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("webサーバの起動に失敗しました。:", err)
	}
}
