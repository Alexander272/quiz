package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type MediaService struct{}

func NewMediaService() *MediaService {
	return &MediaService{}
}

type Media interface {
	SaveFile(file *multipart.FileHeader, dst string) error
	Move(src, dst string) error
	Delete(dst string) error
}

// SaveUploadedFile uploads the form file to specific dst.
func (s *MediaService) SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file. error: %w", err)
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return fmt.Errorf("failed to create path. error: %w", err)
	}

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create file. error: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("failed to copy file data. error: %w", err)
	}
	return nil
}

func (s *MediaService) Move(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return fmt.Errorf("failed to create path. error: %w", err)
	}

	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("failed to move file. error: %w", err)
	}
	return nil
}

func (s *MediaService) Delete(dst string) error {
	if err := os.RemoveAll(dst); err != nil && !strings.Contains(err.Error(), "no such file") {
		return fmt.Errorf("failed to delete file. error: %w", err)
	}
	return nil
}
