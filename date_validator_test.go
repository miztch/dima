package main

import (
	"testing"
)

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		name     string
		dateStr  string
		expected bool
	}{
		{"Valid date", "2023-07-09", true},
		{"Invalid date format", "2023/07/09", false},
		{"Invalid date", "2023-13-45", false},
		{"Empty string", "", false},
		{"Non-date string", "hello", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidDate(tt.dateStr)
			if result != tt.expected {
				t.Errorf("IsValidDate(%s) = %v; want %v", tt.dateStr, result, tt.expected)
			}
		})
	}
}
