//@author Stanislav Polaniev <spolanyev@gmail.com>

package library

import (
	"path/filepath"
	"testing"
)

func TestGetSourceDirectory(t *testing.T) {
	callerFile := "/project/directory/nested/file.go"

	want, err := filepath.Abs(filepath.Dir(filepath.Dir(callerFile)))
	if err != nil {
		t.Fatalf("Failed to build absolute path: %s", err)
	}

	cases := []struct {
		testName   string
		mockCaller func(int) (uintptr, string, int, bool)
		want       string
	}{
		{
			"1",
			func(int) (uintptr, string, int, bool) {
				return 0, callerFile, 0, true
			},
			want,
		},
		{
			"2",
			func(int) (uintptr, string, int, bool) {
				return 0, "", 0, false
			},
			"",
		},
	}

	for _, theCase := range cases {
		got, _ := GetFullPathSourceDirectory(theCase.mockCaller)
		if got != theCase.want {
			t.Errorf("Test %v: source directory == %v, want %v", theCase.testName, got, theCase.want)
		}
	}
}
