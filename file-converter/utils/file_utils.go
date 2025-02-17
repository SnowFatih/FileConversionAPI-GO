package utils
import (
	"os"
	"path/filepath"
	"strings"
)



func GenerateNewFileName(originalName, newExt string) string {
	base := strings.TrimSuffix(originalName, filepath.Ext(originalName))
	return base + newExt
}


func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}