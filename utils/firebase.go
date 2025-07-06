package utils

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseAuth *auth.Client

func InitFirebase() {
	// 先嘗試從環境變數讀取 JSON 字串
	jsonCredentials := os.Getenv("FIREBASE_ADMIN_JSON")

	var opt option.ClientOption
	if jsonCredentials != "" {
		// 環境變數存在，使用它初始化
		opt = option.WithCredentialsJSON([]byte(jsonCredentials))
	} else {
		// 環境變數不存在，改用本地檔案（給本機開發用）
		opt = option.WithCredentialsFile("firebase-admin.json")
	}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase 初始化失敗: %v", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase Auth 初始化失敗: %v", err)
	}
}
