package util_test

import (
	"reflect"
	"testing"

	"github.com/ssulei7/gh-runner-usage/internal/util" // Adjust the import path according to your project structure
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "No Duplicates",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "With Duplicates",
			input:    []string{"a", "b", "a", "c", "b"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty Slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Nil Slice",
			input:    nil,
			expected: []string{},
		},
		{
			name:     "All Duplicates",
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := util.RemoveDuplicates(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
