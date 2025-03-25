package graphs

import (
	"testing"
)

var friends = Graph[string]{
	"alfredo": []string{"heriana", "jesus", "oreana"},
	"heriana": []string{"jesus"},
	"jesus":   []string{"oreana"},
	"oreana":  []string{"gio", "luis"},
	"gio":     []string{"luis"},
}

func TestBFS(t *testing.T) {
	type testCase struct {
		start  string
		target string
		found  bool
	}

	testCases := []testCase{
		{"alfredo", "heriana", true},
		{"alfredo", "jesus", true},
		{"alfredo", "oreana", true},
		{"alfredo", "gio", true},
		{"alfredo", "luis", true},
		{"alfredo", "fernando", false},
	}

	for _, tc := range testCases {
		found := BFS(friends, tc.start, tc.target)

		if tc.found != found {
			t.Errorf("expected %v but got %v for return found value property", tc.found, found)
		}
	}
}
