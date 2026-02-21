package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanSize(t *testing.T) {
	tests := []struct {
		name     string
		bytes    int64
		expected string
	}{
		{"zero", 0, "-"},
		{"negative", -1, "-"},
		{"small MB", 85 * 1024 * 1024, "85.0 MB"},
		{"fractional MB", 85300000, "81.3 MB"},
		{"large MB", 500 * 1024 * 1024, "500.0 MB"},
		{"GB range", 2 * 1024 * 1024 * 1024, "2.00 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HumanSize(tt.bytes)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatDigest(t *testing.T) {
	tests := []struct {
		name     string
		digest   string
		expected string
	}{
		{"empty", "", ""},
		{"sha256 long", "sha256:abcdef123456789012345678", "sha256:abcdef123456"},
		{"sha256 short", "sha256:abc", "sha256:abc"},
		{"other format long", "md5:abcdef1234567890123456", "md5:abcdef123456789"},
		{"other format short", "abc", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDigest(tt.digest)
			assert.Equal(t, tt.expected, result)
		})
	}
}
