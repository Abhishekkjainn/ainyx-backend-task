package models

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday passed this year",
			dob:      now.AddDate(-20, -1, 0), // 20 years and 1 month ago
			expected: 20,
		},
		{
			name:     "Birthday is today",
			dob:      now.AddDate(-20, 0, 0), // Exactly 20 years ago
			expected: 20,
		},
		{
			name:     "Birthday not yet passed this year",
			dob:      now.AddDate(-20, 1, 0), // 20 years ago, but next month
			expected: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			assert.Equal(t, tt.expected, age)
		})
	}
}
