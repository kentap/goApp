package main

import (
	"log"
	"net/http"
	"goApp/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Muxルーターを作成
	r := mux.NewRouter()

	// YouTubeデータを取得するエンドポイントを設定
	r.HandleFunc("/get-youtube-data", handlers.GetYouTubeData).Methods("GET")

	// サーバーの開始
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
