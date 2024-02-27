package helper

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
)

func LognToken(username, password string) string {

	sum := sha256.Sum256([]byte(username + password))
	return fmt.Sprintf("%x", sum)
}

func UploadImage(req *http.Request) string {
	image, header, err := req.FormFile("avatar")
	if err != nil {
		log.Println("❌ Request doesn't contain image", err)
		return ""
	}
	defer image.Close()
	// if header.Size > maxSize {
	// 	log.Println("❌ File size exceeds limit")
	// 	return ""
	// }
	if !isValidFileType(header.Header.Get("Content-Type")) {
		log.Println("❌ Invalid file type")
		return ""
	}
	uploads := "/static/uploads"
	imageURL := filepath.Join(uploads, generateUniqueFilename(header.Filename))
	filePath := filepath.Join(".", imageURL) // Use "." to denote the current directory
	// if filePath[0] != '/' {
	// 	filePath = "" + filePath
	// }
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Error when creating the file", err)
		return ""
	}
	defer file.Close()
	_, err = io.Copy(file, image)
	if err != nil {
		fmt.Println("❌ Error when copying data", err)
		return ""
	}
	return imageURL
}

func generateUniqueFilename(filename string) string {
	ext := filepath.Ext(filename)
	randomName, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	newFilename := randomName.String() + ext
	return newFilename
}
func isValidFileType(contentType string) bool {
	switch contentType {
	case "image/jpeg", "image/png", "image/gif":
		return true
	}
	return false
}
