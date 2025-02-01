package handlers

import (
	"encoding/json"
	"net/http"
	"goApp/services"
	//"goApp/models"
)

// YouTubeデータを取得するハンドラ
func GetYouTubeData(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	query := r.URL.Query().Get("query")
	nums := r.URL.Query().Get("nums")

	if query == "" || nums == "" {
		http.Error(w, "Query and nums parameters are required", http.StatusBadRequest)
		return
	}

	// YouTube APIからデータを取得
	data, err := services.FetchYouTubeData(query, nums)
	if err != nil {
		http.Error(w, "Failed to fetch YouTube data", http.StatusInternalServerError)
		return
	}

	// カスタマイズされたデータを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
