//@author Stanislav Polaniev <spolanyev@gmail.com>

package library

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

func TestFileManipulator(t *testing.T) {
	//create test directory
	testDir := "test_dir"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %s", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to remove test directory: %s", err)
		}
	}(testDir)
	//add files
	testFiles := []string{"file1.txt", "file2.doc", "file3.txt"}
	for _, filename := range testFiles {
		file, err := os.Create(filepath.Join(testDir, filename))
		if err != nil {
			t.Fatalf("Failed to create test file %s: %s", filename, err)
		}
		err = file.Close()
		if err != nil {
			t.Fatalf("Failed to close test file %s: %s", filename, err)
			return
		}
	}
	//create subdirectory
	subDir := filepath.Join(testDir, "sub_dir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %s", err)
	}
	//add files
	subDirFiles := []string{"file4.txt", "file5.pdf"}
	for _, filename := range subDirFiles {
		file, err := os.Create(filepath.Join(subDir, filename))
		if err != nil {
			t.Fatalf("Failed to create test file %s: %s", filename, err)
		}
		err = file.Close()
		if err != nil {
			t.Fatalf("Failed to close test file %s: %s", filename, err)
			return
		}
	}
	//create one more nested directory
	nestedDir := filepath.Join(subDir, "nested_dir")
	err = os.Mkdir(nestedDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create nested: %s", err)
	}
	//add files
	nestedDirFiles := []string{"file6.png", "file7.txt"}
	for _, filename := range nestedDirFiles {
		file, err := os.Create(filepath.Join(nestedDir, filename))
		if err != nil {
			t.Fatalf("Failed to create test file %s: %s", filename, err)
		}

		if filename == "file7.txt" {
			_, err := file.WriteString("line1\nline2\nline3\n\n\n\n")
			if err != nil {
				t.Fatalf("Failed to write to file %s: %s", filename, err)
				return
			}
		}

		err = file.Close()
		if err != nil {
			t.Fatalf("Failed to close test file %s: %s", filename, err)
			return
		}
	}
	//run tests
	t.Run("IsRegularFileExist", func(t *testing.T) {
		cases := []struct {
			file string
			want bool
		}{
			{
				"",
				false,
			},
			{
				"..",
				false,
			},
			{
				filepath.Join(nestedDir, "file6.png"),
				true,
			},
		}

		for _, theCase := range cases {
			fm := &FileManipulator{}
			regularFile := theCase.file
			ok, _ := fm.IsRegularFileExist(regularFile)

			if ok != theCase.want {
				t.Errorf("Regular file exist == %v, want %v", ok, theCase.want)
			}
		}
	})

	t.Run("ReadDirectory", func(t *testing.T) {
		//define callback
		matchedFiles := make([]string, 0)
		getMatchedFiles := func(filename string) {
			matchedFiles = append(matchedFiles, filename)
		}
		//define test cases
		cases := []struct {
			fm      *FileManipulator
			options ReadDirectoryOptions
			want    int
		}{
			{
				&FileManipulator{},
				ReadDirectoryOptions{
					Pattern:     regexp.MustCompile(".txt"),
					Callback:    getMatchedFiles,
					IsRecursive: true,
				},
				4,
			},
			{
				&FileManipulator{},
				ReadDirectoryOptions{
					Pattern:     regexp.MustCompile(".txt"),
					Callback:    getMatchedFiles,
					IsRecursive: false,
				},
				2,
			},
			{
				&FileManipulator{},
				ReadDirectoryOptions{
					Pattern:     nil,
					Callback:    getMatchedFiles,
					IsRecursive: false,
				},
				4,
			},
			{
				&FileManipulator{},
				ReadDirectoryOptions{
					Pattern:     regexp.MustCompile("sub_dir"),
					Callback:    getMatchedFiles,
					IsRecursive: true,
				},
				1,
			},
		}

		for _, theCase := range cases {
			matchedFiles = make([]string, 0)
			theCase.fm.ReadDirectory(testDir, theCase.options)

			if len(matchedFiles) != theCase.want {
				t.Errorf("Matched files == %v, want %v", len(matchedFiles), theCase.want)
			}
		}
	})

	t.Run("Read", func(t *testing.T) {
		cases := []struct {
			file string
			want string
		}{
			{
				"",
				"",
			},
			{
				"..",
				"",
			},
			{
				filepath.Join(nestedDir, "file7.txt"),
				"line1\nline2\nline3\n\n\n\n",
			},
		}

		for _, theCase := range cases {
			fm := &FileManipulator{}
			regularFile := theCase.file
			got, _ := fm.Read(regularFile)

			if got != theCase.want {
				t.Errorf("File content == %v, want %v", got, theCase.want)
			}
		}
	})

	t.Run("GetLineQuantity", func(t *testing.T) {
		cases := []struct {
			file string
			want int
		}{
			{
				"",
				0,
			},
			{
				filepath.Join(nestedDir, "file7.txt"),
				6,
			},
		}

		for _, theCase := range cases {
			fm := &FileManipulator{}
			regularFile := theCase.file
			got, _ := fm.GetLineQuantity(regularFile)

			if got != theCase.want {
				t.Errorf("File line quantity == %v, want %v", got, theCase.want)
			}
		}
	})

	t.Run("GetSlice", func(t *testing.T) {
		cases := []struct {
			file string
			want []string
		}{
			{
				"",
				[]string{},
			},
			{
				filepath.Join(nestedDir, "file7.txt"),
				[]string{"line1", "line2", "line3"},
			},
		}

		for _, theCase := range cases {
			fm := &FileManipulator{}
			regularFile := theCase.file
			got, _ := fm.GetSlice(regularFile)

			if !reflect.DeepEqual(got, theCase.want) {
				t.Errorf("Slice == %v, want %v", got, theCase.want)
			}
		}
	})

}
