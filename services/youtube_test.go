package services

import (
	"testing"
	"goApp/models"   // modelsパッケージを使う場合はインポート
	"fmt"
	"net/http"
	"net/http/httptest"
)

// モック用のYouTube APIレスポンス
func mockYouTubeAPIResponse(query string, nums string) string {
	return `{
		"items": [
			{
				"id": {"videoId": "abcd1234"},
				"snippet": {
					"title": "Golang Tutorial",
					"channelTitle": "GoDev",
					"publishedAt": "2023-01-01T00:00:00Z"
				},
				"statistics": {
					"viewCount": "1000000"
				}
			},
			{
				"id": {"videoId": "efgh5678"},
				"snippet": {
					"title": "Learn Golang",
					"channelTitle": "GoMaster",
					"publishedAt": "2023-02-15T00:00:00Z"
				},
				"statistics": {
					"viewCount": "900000"
				}
			}
		]
	}`
}

// テスト用のHTTPハンドラ
func TestFetchYouTubeData(t *testing.T) {
	// モックのHTTPサーバーを作成
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// モックレスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, mockYouTubeAPIResponse("golang", "2"))
	}))
	defer ts.Close()

	// サーバーのURLをテスト関数に渡す
	apiUrl := ts.URL

	// `FetchYouTubeData`を呼び出して、実際にモックAPIからデータを取得
	data, err := FetchYouTubeData("golang", "2", apiUrl)
	if err != nil {
		t.Fatalf("Error fetching YouTube data: %v", err)
	}

	// テストしたいデータをアサート
	if len(data) != 2 {
		t.Fatalf("Expected 2 items, but got %d", len(data))
	}

	if data[0].VideoID != "abcd1234" {
		t.Errorf("Expected videoId 'abcd1234', but got '%s'", data[0].VideoID)
	}
	if data[1].ViewCount != 900000 {
		t.Errorf("Expected viewCount 900000, but got %d", data[1].ViewCount)
	}
}
