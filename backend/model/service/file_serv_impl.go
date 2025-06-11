package service

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"lampung_trip/model/domain"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type FileServImpl struct {
	File domain.File
	Host string
}

func NewFileServImpl(file domain.File, host string) *FileServImpl {
	return &FileServImpl{File: file, Host: host}
}

func (serv *FileServImpl) UploadFile(files []*multipart.FileHeader) ([]domain.File, error) {
	uploadDir := "assets"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	var uploaded []domain.File

	// Valid extension
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	for _, fileHeader := range files {
		// Validation Exts
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !allowedExts[ext] {
			return nil, fmt.Errorf("file type not allowed: %s", ext)
		}

		// Open file
		src, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}

		// Generate unique file name
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
		dstPath := filepath.Join(uploadDir, fileName)

		// Save to disk
		dst, errDst := os.Create(dstPath)
		if errDst != nil {
			_ = src.Close()
			return nil, fmt.Errorf("failed to create file: %w", errDst)
		}

		if _, err := io.Copy(dst, src); err != nil {
			_ = dst.Close()
			_ = src.Close()
			return nil, fmt.Errorf("failed to save file: %w", err)
		}

		// Close File
		_ = dst.Close()
		_ = src.Close()

		// Create URL FILE
		fileURL := fmt.Sprintf("%s/assets/%s", serv.Host, fileName)

		uploadedFile := domain.File{
			FileName: fileName,
			FileURL:  fileURL,
		}

		uploaded = append(uploaded, uploadedFile)
	}

	return uploaded, nil
}

func (serv *FileServImpl) DeleteFile(fileURL string) error {
	// Parse URL
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return fmt.Errorf("invalid file URL: %w", err)
	}

	// Make sure initial path is "/assets/"
	if !strings.HasPrefix(parsedURL.Path, "/assets/") {
		return fmt.Errorf("file URL must start with /assets/")
	}

	// Get File Name
	fileName := strings.TrimPrefix(parsedURL.Path, "/assets/")
	if fileName == "" {
		return fmt.Errorf("file name is empty in URL")
	}

	// Full path file on disk
	fullPath := filepath.Join("assets", fileName)

	// Check file is exist
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}

	// Delete file
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}
