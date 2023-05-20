//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	stor "dictionary/storage"
	"fmt"
	"path/filepath"
	"runtime"
)

type GetUserFileWords struct {
	FileManipulator *lib.FileManipulator
}

func (cmd *GetUserFileWords) Execute(params map[string]interface{}) dto.ResponseInterface {
	fileName, ok := params["file"].(string)
	if !ok {
		return &dto.ErrorMessage{Message: "Invalid params", From: "GetUserFileWords"}
	}
	fileName = filepath.Base(fileName)

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	directory := filepath.Dir(currentFile)
	fullPathDirectory, err := filepath.Abs(directory)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}

	fullPathFile := filepath.Join(filepath.Dir(fullPathDirectory), stor.PUBLIC_DIR, stor.USER_DATA_DIR, fileName)
	words := cmd.FileManipulator.GetSlice(fullPathFile)
	return &dto.Message{Message: "", From: "getUserFileWords", Data: words, IsError: false}
}
