package utils

import (
	"path/filepath"
	"strings"
)

func IsAllowedImageFile(filename string) (bool, string) {
	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".webp": true,
	}

	ext := filepath.Ext(filename)
	return allowedExtensions[strings.ToLower(ext)], ext
}
