# GolangとWebSocketで作ったリアルタイムなタスク管理アプリ

[![Image from Gyazo](https://i.gyazo.com/fe97603c094419d36683e5d7186eb028.gif)](https://gyazo.com/fe97603c094419d36683e5d7186eb028)

Golangを勉強し始めて2ヶ月ほど経ったので、WebSocketの勉強も兼ねて、リアルタイムタスク管理アプリを作りました。<br>
ユーザーがタスクを追加・完了すると、他のユーザーにもその様子がリアルタイムで表示されるようになっています。
DBとは接続していないので、指定したタスクを完了する部分は、それぞれのタスクにCSSのClassを付けて実現させています。

# 使い方
リポジトリをクローンして、下記を実行する。
```
$ ./golang-websocket-task-app
```

# 参考
- [[ golang ] WebSocketを使ったチャット機能を実装してみる。](http://wild-data-chase.com/index.php/2019/03/20/post-643/)
