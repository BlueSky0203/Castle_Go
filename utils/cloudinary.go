package utils

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(filePath string) (*uploader.UploadResult, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatalf("Failed to init Cloudinary: %v", err)
	}

	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}

	return uploadResult, nil
}
