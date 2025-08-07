package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
	".docx": true,
}

func SaveSingleFile(file *multipart.FileHeader, folder string) (string, error) {
	if err := ensureFolder(folder); err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("file type '%s' not allowed", ext)
	}

	originalName := filepath.Base(file.Filename)
	formatFileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), originalName)
	savePath := filepath.Join(folder, formatFileName)

	return savePath, saveUploadedFile(file, savePath)
}

func SaveMultipleFiles(files []*multipart.FileHeader, folder string) ([]string, error) {
	if err := ensureFolder(folder); err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowedExtensions[ext] {
			continue
		}
		originalName := filepath.Base(file.Filename)
		formatFileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), originalName)
		savePath := filepath.Join(folder, formatFileName)

		if err := saveUploadedFile(file, savePath); err != nil {
			return nil, err
		}
		paths = append(paths, savePath)
	}
	return paths, nil
}

func ensureFolder(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return os.MkdirAll(folder, os.ModePerm)
	}
	return nil
}

func saveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
