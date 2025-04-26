package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// UploadImageFromFile 直接從 multipart 文件上傳到 Cloudinary
func UploadImageFromFile(file *multipart.FileHeader) (*uploader.UploadResult, error) {
	// 從環境變數取得 Cloudinary URL
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("CLOUDINARY_URL is not set")
		return nil, fmt.Errorf("Cloudinary URL not set")
	}

	// 建立 Cloudinary 實例
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Println("Failed to create Cloudinary instance:", err)
		return nil, err
	}

	// 打開上傳的檔案
	imgFile, err := file.Open()
	if err != nil {
		log.Println("Failed to open uploaded file:", err)
		return nil, err
	}
	defer imgFile.Close()

	// 上傳檔案到 Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), imgFile, uploader.UploadParams{})
	if err != nil {
		log.Println("Failed to upload image to Cloudinary:", err)
		return nil, err
	}

	return uploadResult, nil
}
