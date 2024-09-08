package sizecalc

import (
	"testing"
)

func TestFormatSize(t *testing.T) {
	tests := []struct {
		size     int64
		unit     string
		expected string
	}{
		{1024, "KB", "1.00 KB"},
		{1048576, "MB", "1.00 MB"},
		{1073741824, "GB", "1.00 GB"},
		{1073741824, "unknown", "1.00 GB"}, // Default case
	}

	for _, tt := range tests {
		t.Run(tt.unit, func(t *testing.T) {
			result := FormatSize(tt.size, tt.unit)
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}
