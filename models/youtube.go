package models

// YouTube APIのレスポンス構造体
type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
			PublishedAt string `json:"publishedAt"`
		} `json:"snippet"`
		Statistics struct {
			ViewCount   string `json:"viewCount"`
		} `json:"statistics"`
	} `json:"items"`
}

// カスタマイズされたYouTubeレスポンスデータ
type CustomYouTubeResponse struct {
	VideoID    string `json:"videoId"`
	Title      string `json:"title"`
	Channel    string `json:"channelTitle"`
	Published  string `json:"publishedAt"`
	ViewCount  int    `json:"viewCount"`
}
