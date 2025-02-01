package utils

import (
	"log"
)

// エラーハンドリングなどのユーティリティ関数
func HandleError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
