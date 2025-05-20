package repo_test

import (
	"go-hexagonal-architecture/internal/adapter/infra/mongo/repo"
	"testing"
)

func TestCalculateSkip(t *testing.T) {
	tests := []struct {
		name     string
		page     int64
		size     int64
		expected int64
	}{
		{
			name:     "valid page and size",
			page:     2,
			size:     10,
			expected: 10,
		},
		{
			name:     "negative page",
			page:     -1,
			size:     10,
			expected: 0,
		},
		{
			name:     "zero size",
			page:     1,
			size:     0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.CalculateSkip(tt.page, tt.size)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
