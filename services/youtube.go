package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"goApp/models"
	"strconv"
	"sort"
	"io/ioutil"
)

// YouTube APIからデータを取得し、カスタマイズされたデータを返す関数
func FetchYouTubeData(query string, nums string) ([]models.CustomYouTubeResponse, error) {
	apiKey := "AIzaSyDcxIvYnkgXr1VyZkYlkFYTUw9cNhUrYuE"
	apiUrl := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&order=viewCount&maxResults=%s&key=%s", query, nums, apiKey)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var youtubeResp models.YouTubeResponse
	err = json.Unmarshal(body, &youtubeResp)
	if err != nil {
		return nil, err
	}

	// カスタマイズされたデータを作成
	var customData []models.CustomYouTubeResponse
	for _, item := range youtubeResp.Items {
		viewCount, _ := strconv.Atoi(item.Statistics.ViewCount) // 再生回数を整数に変換
		customData = append(customData, models.CustomYouTubeResponse{
			VideoID:   item.ID.VideoID,
			Title:     item.Snippet.Title,
			Channel:   item.Snippet.ChannelTitle,
			Published: item.Snippet.PublishedAt,
			ViewCount: viewCount,
		})
	}

	// 再生回数順に並び替え
	sort.Slice(customData, func(i, j int) bool {
		return customData[i].ViewCount > customData[j].ViewCount
	})

	return customData, nil
}
