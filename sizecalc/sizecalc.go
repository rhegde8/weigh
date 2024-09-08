package sizecalc

import (
	"fmt"
	"os"
	"path/filepath"
)

// CalculateSize returns the size of a file or directory in bytes.
func CalculateSize(path string) (int64, error) {
	var totalSize int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to walk through path: %w", err)
	}

	return totalSize, nil
}

// FormatSize converts bytes into the specified format (KB, MB, GB).
func FormatSize(size int64, unit string) string {
	switch unit {
	case "KB":
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	case "MB":
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	case "GB":
		return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024)) // Default to GB
	}
}
