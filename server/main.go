package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	StartBackendServer()
}

func StartBackendServer() {
	// buildフォルダを公開
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../client/build"))))
	// rootディレクトリにアクセスしたときに相対パスで指定した静的サイトをホストする
	// 相対ディレクトリを指定するために client/package.jsonに　"homepage": "./"　を追記する

	// 上記の"homepage": "./"を記述してないと下の記述でしか動かない。他のパスだとバグる
	// http.Handle("/", http.FileServer(http.Dir("./build")))

	fmt.Println("Server Started Port 8080")

	http.HandleFunc("/api/ping", pingResponse)
	http.HandleFunc("/api", samplePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Port:" + port)
	http.ListenAndServe(":"+port, nil) // Herokuで実装する場合は環境変数からPORTを取得する
	//http.ListenAndServe(":8080", nil)
}

func pingResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong2")
}

func samplePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "おはよう世界")
}
